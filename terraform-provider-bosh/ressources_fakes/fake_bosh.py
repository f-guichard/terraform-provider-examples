# -*- coding: UTF-8 -*-

#Othello.java style : single file program

import os
from flask import Flask
from flask import jsonify
from flask import request

# Global variables section (ouaip Alice...)
PORT = 30002 #Tu es dans ton container, back to back service
CONTROLLER_VERSION = "v1"
_CONTROLLER_NAME = "Bosh Director"
# Better use a regex matcher class
_STEMCELLS_URL = "/"+CONTROLLER_VERSION+"/api-director/stemcells"
_STEMCELL_ID = "/"+CONTROLLER_VERSION+"/api-director/stemcells/<cid>"
_RELEASES_URL = "/"+CONTROLLER_VERSION+"/api-director/releases"
_RELEASE_ID = "/"+CONTROLLER_VERSION+"/api-director/releases/<cid>"
_DEPLOYMENTS_URL = "/"+CONTROLLER_VERSION+"/api-director/deployments"
_DEPLOYMENT_ID = "/"+CONTROLLER_VERSION+"/api-director/deployments/<cid>"


_HELPER_RESPONSE = {
    _CONTROLLER_NAME: CONTROLLER_VERSION,
    "GET ["+_STEMCELLS_URL+","+_RELEASES_URL+","+_DEPLOYMENTS_URL+"]" : {
        "method": "GET",
        "parameters": "",
        "code retour": "200",
        "description": "Lister tous les [stemcells,releases,deploiements]"
    },
    "GET ["+_STEMCELL_ID+","+_RELEASE_ID+","+_DEPLOYMENT_ID+"]" : {
        "method": "GET",
        "parameters": "un identifiant de [stemcell,release,deploiement]",
        "code retour": "200",
        "description": "Lister le detail d'un [stemcell,release,deploiement]"
    },
    "POST ["+_STEMCELLS_URL+","+_RELEASES_URL+","+_DEPLOYMENTS_URL+"]" : {
        "method": "POST",
        "parameters": "json body",
        "code retour": "201",
        "description": "Creer un [stemcell,release,deploiement]"
    },
    "PATCH ["+_STEMCELL_ID+","+_RELEASE_ID+","+_DEPLOYMENT_ID+"]" : {
        "method": "PATCH",
        "parameters": "json body",
        "code retour": "200",
        "description": "Mettre a jour la description d'un [stemcell,release,deploiement]"
    },
    "DELETE ["+_STEMCELL_ID+","+_RELEASE_ID+","+_DEPLOYMENT_ID+"]" : {
        "method": "DELETE",
        "parameters": "un identifiant de [stemcell,release,deploiement]",
        "code retour": "200",
        "description": "Supprimer un [stemcell,release,deploiement]"
    }
}

#Structure de persistence (ouaip Bob...)
ramDic = {}

## Creation & configuration du controleur Flask
app = Flask(__name__)

@app.route('/')
def index():
    return 'WORKING'

@app.route('/help')
def help():
    return jsonify(_HELPER_RESPONSE)

@app.route(_STEMCELLS_URL, methods=['GET'])
@app.route(_RELEASES_URL, methods=['GET'])
@app.route(_DEPLOYMENTS_URL, methods=['GET'])
def list_stemcells():
    #PEP 448
    response = jsonify(*ramDic)
    response.status_code = 200
    return response

@app.route(_STEMCELL_ID, methods=['GET'])
@app.route(_RELEASE_ID, methods=['GET'])
@app.route(_DEPLOYMENT_ID, methods=['GET'])
def list_stemcell(cid):
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.get(cid))
    response.status_code = 200
    return response

@app.route(_STEMCELLS_URL, methods=['POST'])
@app.route(_RELEASES_URL, methods=['POST'])
@app.route(_DEPLOYMENTS_URL, methods=['POST'])
def create_stemcells():
    body = request.get_json(force=True)
    #Ouchhhhhhhhhh
    ramDic[str(len(ramDic))] = body
    response = jsonify({'cid':str(len(ramDic)-1)},{"stemcell":ramDic.get(str(len(ramDic)-1))})
    response.status_code = 201
    return response

@app.route(_STEMCELL_ID, methods=['PATCH'])
@app.route(_RELEASE_ID, methods=['PATCH'])
@app.route(_DEPLOYMENT_ID, methods=['PATCH'])
def patch_stemcells():
    response = jsonify('NOT IMPLEMENTED YET')
    response.status_code = 200
    return response

@app.route(_STEMCELL_ID, methods=['DELETE'])
@app.route(_RELEASE_ID, methods=['DELETE'])
@app.route(_DEPLOYMENT_ID, methods=['DELETE'])
def delete_stemcellss(cid):
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.pop(cid))
    response.status_code = 200
    return response

app.debug = True
app.run(host='0.0.0.0', port=PORT)