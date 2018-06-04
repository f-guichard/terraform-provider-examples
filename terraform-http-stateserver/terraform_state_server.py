# -*- coding: UTF-8 -*-

import os
import json

from flask import Flask
from flask import jsonify
from flask import request

#See https://www.terraform.io/docs/backends/types/http.html
PORT = 40000 #Tu es dans ton container, back to back service
CONTROLLER_VERSION = "v1"
_CONTROLLER_NAME = "Terraform State Controller"
_TFSTATES_URL = "/"+CONTROLLER_VERSION+"/states"
_TFSTATE_ID = "/"+CONTROLLER_VERSION+"/states/<stateid>"

#Basic Auth Forced
#Set to False if not desired
_HTTP_AUTH_REQUIRED = True

_HELPER_RESPONSE = {
    _CONTROLLER_NAME: CONTROLLER_VERSION,
    "GET "+_TFSTATES_URL : {
        "method": "GET",
        "parameters": "",
        "code retour": "200",
        "description": "Lister tous les states terraform"
    },
    "GET "+_TFSTATE_ID : {
        "method": "GET",
        "parameters": "un identifiant de state terraform",
        "code retour": "200",
        "description": "Lister le detail d'un state terraform"
    },
    "POST "+_TFSTATE_ID : {
        "method": "POST",
        "parameters": "json body like see any terraform.tfstate file",
        "code retour": "201",
        "description": "Creer un state terraform"
    },
    "DELETE "+_TFSTATE_ID : {
        "method": "DELETE",
        "parameters": "un identifiant d'un state terraform",
        "code retour": "200",
        "description": "Supprimer un state terraform"
    }
}

######### Rework all that
#Extraire le pass d'un user
def getPass(user, dbpasswd, tfstateid):
    return dbpasswd.get(tfstateid)["password"]

##Lire le contenu de la database des utilisateurs autorises
def loaddb():
    try:
        with open('db/users.db', 'r') as usersdb:
            return json.load(usersdb)
    except Exception as ex:
        print('Erreur : {}'.format(ex))
        exit(1)
#########

##Ecrire une methode qui persiste sur le fileystem
def persiststate(tfstatestruct):
    try:
        with open('tfstatefile.json', 'w') as tfstatefile:
            json.dump(tfstatestruct, tfstatefile)
    except Exception as ex:
        print('Erreur : {}'.format(ex))
        exit(1)

#Structure de stockage des users/log
credz = loaddb()

#Structure de persistence (ouaip Bob...)
ramDic = {"example":{"version": 3, "terraform_version": "0.11.7", "serial": 1, \
  "lineage": "b457bdd4-2cf4-3c06-eb8f-f422c3a99fe3", "modules": \
[ {"path": ["root"], "outputs": {}, "resources": {}, "depends_on": [] } ]\
}
}

## Creation & configuration du controleur Flask
app = Flask(__name__)

@app.route('/')
def index():
    return 'WORKING'

@app.route('/help')
def help():
    return jsonify(_HELPER_RESPONSE)

@app.route(_TFSTATES_URL, methods=['GET'])
def list_states():
    #PEP 448
    response = jsonify(*ramDic)
    response.status_code = 200
    return response

@app.route(_TFSTATE_ID, methods=['GET'])
def list_state(stateid):
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.get(stateid))
    response.status_code = 404 
    if ramDic.get(stateid) != None:
        response.status_code = 200
    return response

@app.route(_TFSTATE_ID, methods=['POST'])
def create_state(stateid):
    body = request.get_json(force=True)
    ###Ouch bis ! check for framework || decorated function
    if _HTTP_AUTH_REQUIRED:
        user = request.authorization.username
        passwd = request.authorization.password
        if passwd != getPass(user, credz, stateid):
            response = jsonify("Need proper credz")
            response.status_code = 401
            return response
    #Ouchhhhhhhhhh
    ramDic[stateid] = body
    response = jsonify({'name':stateid},{"tfstate":ramDic.get(stateid)})
    persiststate(ramDic)
    response.status_code = 201
    return response

@app.route(_TFSTATE_ID, methods=['DELETE'])
def delete_flows(stateid):
    ###Ouch bis ! check for framework || decorated function
    if _HTTP_AUTH_REQUIRED:
        user = request.authorization.username
        passwd = request.authorization.password
        if passwd != getPass(user, credz, stateid):
            response = jsonify("Need proper credz")
            response.status_code = 401
            return response
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.pop(stateid))
    persiststate(ramDic)
    response.status_code = 200
    return response

app.debug = True
app.run(host='0.0.0.0', port=PORT)
