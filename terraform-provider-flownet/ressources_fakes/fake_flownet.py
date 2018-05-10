# -*- coding: UTF-8 -*-

#Othello.java style : single file program

import os
from flask import Flask
from flask import jsonify
from flask import request

# Global variables section (ouaip Alice...)
PORT = 30000 #Tu es dans ton container, back to back service
CONTROLLER_VERSION = "v1"
_CONTROLLER_NAME = "Flow Controller"
_FLOWS_URL = "/"+CONTROLLER_VERSION+"/flows"
_FLOW_ID = "/"+CONTROLLER_VERSION+"/flows/<flowid>"

_HELPER_RESPONSE = {
    _CONTROLLER_NAME: CONTROLLER_VERSION,
    "GET "+_FLOWS_URL : {
        "method": "GET",
        "parameters": "",
        "code retour": "200",
        "description": "Lister tous les identifiants de flows"
    },
    "GET "+_FLOW_ID : {
        "method": "GET",
        "parameters": "un identifiant de flow",
        "code retour": "200",
        "description": "Lister le detail d'un flow"
    },
    "POST "+_FLOWS_URL : {
        "method": "POST",
        "parameters": "json body like {'CIDR':'CIDR', 'PORT':'PORT', 'DESCRIPTION':'DESCRIPTION'}",
        "code retour": "201",
        "description": "Creer un flow"
    },
    "PATCH "+_FLOW_ID : {
        "method": "PATCH",
        "parameters": "json body like : {flowid : 'DESCRIPTION':'DESCRIPTION'}",
        "code retour": "200",
        "description": "Mettre a jour la description d'un flow"
    },
    "DELETE "+_FLOW_ID : {
        "method": "DELETE",
        "parameters": "un identifiant de flow",
        "code retour": "200",
        "description": "Supprimer un flow"
    }
}

#Structure de persistence (ouaip Bob...)
ramDic = {"0":{"CIDR":"224.0.0.0/24","PORT":"5894-6500","DESCRIPTION":"1st Flow init (vector/salt)"},\
"1":{"CIDR":"81.26.36.46","PORT":"589","DESCRIPTION":"2nd Flow"}}

## Creation & configuration du controleur Flask
app = Flask(__name__)

@app.route('/')
def index():
    return 'WORKING'

@app.route('/help')
def help():
    return jsonify(_HELPER_RESPONSE)

@app.route(_FLOWS_URL, methods=['GET'])
def list_flows():
    #PEP 448
    response = jsonify(*ramDic)
    response.status_code = 200
    return response

@app.route(_FLOW_ID, methods=['GET'])
def list_flow(flowid):
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.get(flowid))
    response.status_code = 200
    return response

@app.route(_FLOWS_URL, methods=['POST'])
def create_flows():
    body = request.get_json(force=True)
    #Ouchhhhhhhhhh
    ramDic[str(len(ramDic))] = body
    response = jsonify({'flowid':str(len(ramDic)-1)},{"flow":ramDic.get(str(len(ramDic)-1))})
    response.status_code = 201
    return response

@app.route(_FLOW_ID, methods=['PATCH'])
def patch_flows():
    response = jsonify('NOT IMPLEMENTED YET')
    response.status_code = 200
    return response

@app.route(_FLOW_ID, methods=['DELETE'])
def delete_flows(flowid):
    #Ouchhhhhhhhhh
    response = jsonify(ramDic.pop(flowid))
    response.status_code = 200
    return response

app.debug = True
app.run(host='0.0.0.0', port=PORT)