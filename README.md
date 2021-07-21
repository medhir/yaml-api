# API Server Exercise

## Running the Application
You must have an installation of Go in order to run the application. Run the following command within the application's directory:
```sh
go run app.go
```
This command will start the server and accept traffic on port 1111. 

## Using the API 

The API can be accessed from `localhost:1111` and includes `GET` and `POST` http methods to the `/metadata` resource.

### `POST /metadata`

`POST` requests to the `/metadata` resource must include YAML as text in the request body. The YAML must be properly formatted and contain all the attributes listed in the below example: 

```yaml
title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
 ### Interesting Title
 Some application content, and description
```

A successful request will store and index the metadata. 

### `GET /metadata`

`GET` requests to `/metadata` will return any stored metadata by running a search against specific attributes using query parameters. The attributes you can query against include: 

- `title`
- `version`
- `maintainer_name`
- `maintainer_email`
- `company`
- `website`
- `source`
- `license`
- `description`

### Examples 
To find all the metadata where the source includes `github.com`, you could write a query such as `/metadata?source=github.com`.

To search descriptions, you could write a query such as `/metadata?description=some%20application%20content`.

You can also find all metadata that matches multiple fields, such as `/metadata?license=Apache-2.0&title=valid`.