// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kazan hackathon API",
    "title": "Kazan API",
    "version": "0.3.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/",
  "paths": {
    "/order": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "order"
        ],
        "summary": "create order",
        "parameters": [
          {
            "description": "The GitHub API url to call",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/OrderRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/OrderCreateResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    },
    "/order/{id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "order"
        ],
        "summary": "get order",
        "parameters": [
          {
            "type": "string",
            "description": "The order ID.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/OrderResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "order"
        ],
        "summary": "cancel order",
        "parameters": [
          {
            "type": "string",
            "description": "The order ID.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "Canceled",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    },
    "/ticket/{id}/route": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "route",
          "ticket"
        ],
        "summary": "get route by ticket number",
        "parameters": [
          {
            "type": "string",
            "description": "The ticket ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/RouteResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CafeDishResponse": {
      "type": "object",
      "required": [
        "id",
        "name",
        "price",
        "image_url"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "image_url": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "price": {
          "type": "integer"
        }
      }
    },
    "CafeResponse": {
      "type": "object",
      "required": [
        "id",
        "name",
        "city_id",
        "positions"
      ],
      "properties": {
        "city_id": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "positions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/CafeDishResponse"
          }
        }
      }
    },
    "OrderCreateResponse": {
      "type": "object",
      "required": [
        "id",
        "payment_url",
        "status",
        "positions"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "payment_url": {
          "type": "string"
        },
        "positions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/OrderItem"
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "OrderItem": {
      "type": "object",
      "required": [
        "id",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "OrderRequest": {
      "type": "object",
      "required": [
        "user_id",
        "order"
      ],
      "properties": {
        "order": {
          "type": "object",
          "required": [
            "cafe_id",
            "positions"
          ],
          "properties": {
            "cafe_id": {
              "type": "string"
            },
            "positions": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/OrderItem"
              }
            }
          }
        },
        "user_id": {
          "type": "string"
        }
      }
    },
    "OrderResponse": {
      "type": "object",
      "required": [
        "id",
        "status",
        "positions"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "positions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/OrderItem"
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "RouteResponse": {
      "type": "object",
      "required": [
        "train_number",
        "stops"
      ],
      "properties": {
        "stops": {
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "city_id",
              "name",
              "date_time",
              "duration",
              "cafes"
            ],
            "properties": {
              "cafes": {
                "type": "array",
                "items": {
                  "$ref": "#/definitions/CafeResponse"
                }
              },
              "city_id": {
                "type": "string"
              },
              "date_time": {
                "type": "string",
                "format": "date-time"
              },
              "duration": {
                "type": "integer"
              },
              "name": {
                "type": "string"
              }
            }
          }
        },
        "train_number": {
          "type": "string"
        }
      }
    },
    "StatusResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Food order",
      "name": "order"
    },
    {
      "description": "Ticket",
      "name": "ticket"
    },
    {
      "description": "Railroad route",
      "name": "route"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kazan hackathon API",
    "title": "Kazan API",
    "version": "0.3.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/",
  "paths": {
    "/order": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "order"
        ],
        "summary": "create order",
        "parameters": [
          {
            "description": "The GitHub API url to call",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/OrderRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/OrderCreateResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    },
    "/order/{id}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "order"
        ],
        "summary": "get order",
        "parameters": [
          {
            "type": "string",
            "description": "The order ID.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/OrderResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "order"
        ],
        "summary": "cancel order",
        "parameters": [
          {
            "type": "string",
            "description": "The order ID.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "Canceled",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    },
    "/ticket/{id}/route": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "route",
          "ticket"
        ],
        "summary": "get route by ticket number",
        "parameters": [
          {
            "type": "string",
            "description": "The ticket ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/RouteResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CafeDishResponse": {
      "type": "object",
      "required": [
        "id",
        "name",
        "price",
        "image_url"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "image_url": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "price": {
          "type": "integer"
        }
      }
    },
    "CafeResponse": {
      "type": "object",
      "required": [
        "id",
        "name",
        "city_id",
        "positions"
      ],
      "properties": {
        "city_id": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "positions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/CafeDishResponse"
          }
        }
      }
    },
    "OrderCreateResponse": {
      "type": "object",
      "required": [
        "id",
        "payment_url",
        "status",
        "positions"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "payment_url": {
          "type": "string"
        },
        "positions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/OrderItem"
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "OrderItem": {
      "type": "object",
      "required": [
        "id",
        "amount"
      ],
      "properties": {
        "amount": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "OrderRequest": {
      "type": "object",
      "required": [
        "user_id",
        "order"
      ],
      "properties": {
        "order": {
          "type": "object",
          "required": [
            "cafe_id",
            "positions"
          ],
          "properties": {
            "cafe_id": {
              "type": "string"
            },
            "positions": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/OrderItem"
              }
            }
          }
        },
        "user_id": {
          "type": "string"
        }
      }
    },
    "OrderResponse": {
      "type": "object",
      "required": [
        "id",
        "status",
        "positions"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "positions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/OrderItem"
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "RouteResponse": {
      "type": "object",
      "required": [
        "train_number",
        "stops"
      ],
      "properties": {
        "stops": {
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "city_id",
              "name",
              "date_time",
              "duration",
              "cafes"
            ],
            "properties": {
              "cafes": {
                "type": "array",
                "items": {
                  "$ref": "#/definitions/CafeResponse"
                }
              },
              "city_id": {
                "type": "string"
              },
              "date_time": {
                "type": "string",
                "format": "date-time"
              },
              "duration": {
                "type": "integer"
              },
              "name": {
                "type": "string"
              }
            }
          }
        },
        "train_number": {
          "type": "string"
        }
      }
    },
    "StatusResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Food order",
      "name": "order"
    },
    {
      "description": "Ticket",
      "name": "ticket"
    },
    {
      "description": "Railroad route",
      "name": "route"
    }
  ]
}`))
}
