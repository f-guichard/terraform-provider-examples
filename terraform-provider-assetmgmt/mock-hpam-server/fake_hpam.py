# -*- coding: UTF-8 -*-

#Othello.java style : single file program

import os
import time
from flask import Flask
from flask import jsonify
from flask import request

# Global variables section
_CREATE_DELAY = 2
PORT = 30026 # Affectation port a updater pour CloudFoundry
CONTROLLER_VERSION = "v1"
_CONTROLLER_NAME = "Asset Mgmt Controller"
_26E_URL = "/"+CONTROLLER_VERSION+"/26e"
_26E_ID = "/"+CONTROLLER_VERSION+"/26e/<id>"

_HELPER_RESPONSE = {
    _CONTROLLER_NAME: CONTROLLER_VERSION,
    "GET "+_26E_URL : {
        "method": "GET",
        "parameters": "",
        "code retour": "200"
    },
    "GET "+_26E_ID : {
        "method": "GET",
        "parameters": "un identifiant de vips",
        "code retour": "200"
    },
    "POST "+_26E_URL : {
        "method": "POST",
        "parameters": "json body like {}",
        "code retour": "201"
    },
    "PATCH "+_26E_ID : {
        "method": "PATCH",
        "parameters": "json body like : {vipid : 'DESCRIPTION':'DESCRIPTION'}",
        "code retour": "200"
    },
    "DELETE "+_26E_ID : {
        "method": "DELETE",
        "parameters": "un identifiant de vip",
        "code retour": "200"
    }
}

ramDic = {}

app = Flask(__name__)

@app.route('/')
def index():
    return 'WORKING'

@app.route('/help')
def help():
    return jsonify(_HELPER_RESPONSE)

@app.route(_26E_URL, methods=['GET'])
def list_assets():
    #PEP 448
    response = jsonify(*ramDic)
    response.status_code = 200
    return response

@app.route(_26E_ID, methods=['GET'])
def list_asset(id):
    response = jsonify(ramDic.get(id))
    response.status_code = 200
    return response

@app.route(_26E_URL, methods=['POST'])
def create_assets():
    body = request.get_json(force=True)
    ramDic[str(len(ramDic))] = body
    response = jsonify({'id':str(len(ramDic)-1)},{"obj":ramDic.get(str(len(ramDic)-1))})
    response.status_code = 201
    time.sleep(_CREATE_DELAY)
    return response

@app.route(_26E_ID, methods=['PATCH'])
def patch_assets():
    response = jsonify('NOT IMPLEMENTED YET')
    response.status_code = 200
    return response

@app.route(_26E_ID, methods=['DELETE'])
def delete_assets(id):
    response = jsonify(ramDic.pop(id))
    response.status_code = 200
    return response

app.debug = True
app.run(host='0.0.0.0', port=PORT)
