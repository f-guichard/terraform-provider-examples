# -*- coding: UTF-8 -*-

#Othello.java style : single file program

import os
from flask import Flask
from flask import jsonify
from flask import request

# Global variables section (ouaip Alice...)
PORT = 30005 #Tu es dans ton container, back to back service
CONTROLLER_VERSION = "v1"
_CONTROLLER_NAME = "Netlb Controller"
_VIPS_URL = "/"+CONTROLLER_VERSION+"/dns"
_VIP_ID = "/"+CONTROLLER_VERSION+"/dns/<id>"

_HELPER_RESPONSE = {
    _CONTROLLER_NAME: CONTROLLER_VERSION,
    "GET "+_VIPS_URL : {
        "method": "GET",
        "parameters": "",
        "code retour": "200",
        "description": "Lister tous les identifiants de vips"
    },
    "GET "+_VIP_ID : {
        "method": "GET",
        "parameters": "un identifiant de vips",
        "code retour": "200",
        "description": "Lister le detail d'une vip"
    },
    "POST "+_VIPS_URL : {
        "method": "POST",
        "parameters": "json body like {}",
        "code retour": "201",
        "description": "Creer une vip"
    },
    "PATCH "+_VIP_ID : {
        "method": "PATCH",
        "parameters": "json body like : {vipid : 'DESCRIPTION':'DESCRIPTION'}",
        "code retour": "200",
        "description": "Mettre a jour la description d'une vip"
    },
    "DELETE "+_VIP_ID : {
        "method": "DELETE",
        "parameters": "un identifiant de vip",
        "code retour": "200",
        "description": "Supprimer un vip"
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

@app.route(_VIPS_URL, methods=['GET'])
def list_flows():
    #PEP 448
    response = jsonify(*ramDic)
    response.status_code = 200
    return response

@app.route(_VIP_ID, methods=['GET'])
def list_flow(id):
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.get(id))
    response.status_code = 200
    return response

@app.route(_VIPS_URL, methods=['POST'])
def create_flows():
    body = request.get_json(force=True)
    #Ouchhhhhhhhhh
    ramDic[str(len(ramDic))] = body
    response = jsonify({'id':str(len(ramDic)-1)},{"obj":ramDic.get(str(len(ramDic)-1))})
    response.status_code = 201
    return response

@app.route(_VIP_ID, methods=['PATCH'])
def patch_flows():
    response = jsonify('NOT IMPLEMENTED YET')
    response.status_code = 200
    return response

@app.route(_VIP_ID, methods=['DELETE'])
def delete_flows(id):
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.pop(id))
    response.status_code = 200
    return response

app.debug = True
app.run(host='0.0.0.0', port=PORT)