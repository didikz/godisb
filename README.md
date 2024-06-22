# GODISB - DISBURSEMENT API

A simple REST API endpoint written with Go for disbursed money to the beneficiary.

## Setup & Requirements

TBD

## Installation

TBD

## Development

TBD

## API Features

+ [ ] Error Handling
+ [ ] Bank Partner Validation
+ [ ] Idempotency request to prevent double transactions. If request hit using same request id, it will return created transaction instead
+ [ ] Request signature

## API DOCUMENTATION

+ Endpoint: `{baseurl}/api/v1/disbursements`
+ Method: `POST`
+ Headers:
  + `Content-Type`: `application/json`
  + `X-Idempotency-Key`: unique identifier
  + `X-Signature`: Request signature
+ Request Payload:
  + `bank` (string, mandatory): Available beneficiary bank codes. Possible values are `bca`, `bni`, `mandiri`, `bri`
  + `account_number` (string, mandatory): Beneficiary bank account number
  + `amount` (int64, mandatory): Amount of disbursement money
  + `remark` (string, optional, max: 15): Disbursement remark

  sample payload:

    ```json
    {
        "bank": "bca",
        "account_number": "12345679",
        "amount": 10000,
        "remark": "lorem ipsum"
    }
    ```

+ Success Response (`200`)

  Response Object
  + `id` (int64, mandatory): ID of disbursement
  + `bank` (string, mandatory): Beneficiary bank code
  + `account_number` (string, mandatory): Beneficiary bank account number
  + `beneficiary_name` (string, optional): Beneficiary account name holder. Could be empty.
  + `amount` (int64, mandatory): Amount of disbursement
  + `remark` (string, optional): Remark of disbursement
  + `status` (string, mandatory): Status of disbursement. Possible values are:

    + `SUCCESS`: Disbursement succeeded
    + `FAILED`: Disbursement failed
  
  + `failed_notes` (string, optional): Disbursement failed status note codes to give more context why transation failed. Possible values are:

    + `BANK_PARTNER_ERROR`: Error in bank partner when requesting transfer
    + `CANT_RECEIVE_FUNDS`: Error in beneficiary cannot receive funds from bank partner.

  + `created_at` (string, mandatory): Disbursement created date time
  + `failed_at` (string, optional): Disbursement failed date time. Will empty when transaction is succeeded
  + `completed_at` (string, optional): Disbursement completed date time. Will empty when transaction is failed.
  
  Sample payload of **Successful Disbursement**

  ```json
  {
    "id": 12345678,
    "bank": "bca",
    "account_number": "12345679",
    "beneficiary_name": "John Doe",
    "amount": 10000,
    "remark": "lorem ipsum",
    "status": "SUCCESS",
    "failed_notes": "",
    "created_at": "2024-10-10 10:10:10",
    "failed_at": "",
    "completed_at": "2024-10-10 10:10:10"
  }
  ```

  Sample payload of **Failed Disbursement**

  ```json
  {
    "id": 12345678,
    "bank": "bca",
    "account_number": "12345679",
    "beneficiary_name": "John Doe",
    "amount": 10000,
    "remark": "lorem ipsum",
    "status": "FAILED",
    "failed_notes": "BANK_PARTNER_ERROR",
    "created_at": "2024-10-10 10:10:10",
    "failed_at": "",
    "completed_at": ""
  }
  ```

+ Failed Response (`4xx`, `5xx`)

  TBD
