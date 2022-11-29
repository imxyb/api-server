// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "xueyouchen",
            "email": "xueyou@starboardventures.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/complieversions": {
            "get": {
                "description": "list compile version",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.CompileVersionList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/contract/{address}": {
            "get": {
                "description": "Get contract detail",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Contract"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/contract/{address}/internal_txns": {
            "get": {
                "description": "List contract's internal transactions",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "name": "l",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "o",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/busi.EVMInternalTX"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/contract/{address}/is_verify": {
            "get": {
                "description": "get contract is verify",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.ContractIsVerify"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/contract/{address}/txns": {
            "get": {
                "description": "List contract's transactions",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "name": "l",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "o",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/busi.EVMTransaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/contracts": {
            "get": {
                "description": "List contracts",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "name": "l",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "o",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Contract"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/contractverify/{address}": {
            "post": {
                "description": "submit contract verify",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "description": "SubmitContractVerifyRequest",
                        "name": "SubmitContractVerifyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.SubmitContractVerifyRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/busi.EVMContractVerify"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/contractverify/{id}": {
            "get": {
                "description": "check contract verify",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "DATA-INFRA-API-External-V1"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.ContractVerify"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseWithRequestId"
                        }
                    }
                }
            }
        },
        "/api/v1/ping": {
            "get": {
                "description": "Healthy examination",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "Sys"
                ],
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error:...",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "busi.EVMContractVerify": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "compiler_type": {
                    "type": "integer"
                },
                "compiler_version": {
                    "type": "string"
                },
                "contract_name": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "license_type": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "busi.EVMInternalTX": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "parent_hash": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "busi.EVMTransaction": {
            "type": "object",
            "properties": {
                "block_hash": {
                    "type": "string"
                },
                "block_number": {
                    "type": "integer"
                },
                "chain_id": {
                    "type": "integer"
                },
                "from": {
                    "type": "string"
                },
                "gas": {
                    "type": "integer"
                },
                "gas_limit": {
                    "type": "integer"
                },
                "hash": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "input": {
                    "type": "string"
                },
                "max_fee_per_gas": {
                    "type": "string"
                },
                "max_priority_fee_per_gas": {
                    "type": "string"
                },
                "nonce": {
                    "type": "integer"
                },
                "r": {
                    "type": "string"
                },
                "s": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                },
                "transaction_index": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                },
                "v": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "core.CompileVersion": {
            "type": "object",
            "properties": {
                "long_version": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "core.CompileVersionList": {
            "type": "object",
            "properties": {
                "versions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/core.CompileVersion"
                    }
                }
            }
        },
        "core.Contract": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "balance": {
                    "type": "string"
                },
                "compilerType": {
                    "type": "integer"
                },
                "compilerVersion": {
                    "type": "string"
                },
                "filecoinAddress": {
                    "type": "string"
                },
                "license": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "txns": {
                    "type": "integer"
                },
                "verified": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "core.ContractIsVerify": {
            "type": "object",
            "properties": {
                "is_verify": {
                    "type": "boolean"
                }
            }
        },
        "core.ContractVerify": {
            "type": "object",
            "properties": {
                "abi": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "bytecode": {
                    "type": "string"
                },
                "compiler_type": {
                    "type": "integer"
                },
                "compiler_version": {
                    "type": "string"
                },
                "contract_name": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "err_msg": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "license_type": {
                    "type": "string"
                },
                "source_codes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/core.SourceCode"
                    }
                },
                "status": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "core.SourceCode": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "file_name": {
                    "type": "string"
                }
            }
        },
        "core.SourceCodePart": {
            "type": "object",
            "properties": {
                "filename": {
                    "type": "string"
                },
                "source_code_url": {
                    "type": "string"
                }
            }
        },
        "core.SubmitContractVerifyRequest": {
            "type": "object",
            "required": [
                "compiler_type",
                "compiler_version",
                "license_type",
                "runs"
            ],
            "properties": {
                "compiler_type": {
                    "description": "TODO need support more CompilerType, only implement single file now.",
                    "type": "integer",
                    "enum": [
                        1,
                        2,
                        3
                    ]
                },
                "compiler_version": {
                    "type": "string"
                },
                "evm_version": {
                    "type": "string"
                },
                "is_optimization": {
                    "type": "boolean"
                },
                "license_type": {
                    "type": "string"
                },
                "runs": {
                    "description": "TODO need more field support more CompilerType",
                    "type": "integer"
                },
                "source_code": {
                    "type": "string"
                },
                "source_code_parts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/core.SourceCodePart"
                    }
                }
            }
        },
        "utils.ResponseWithRequestId": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "block-explorer-api.spacescope.io",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "api server",
	Description:      "spacescope block explorer api server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
