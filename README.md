# gorma-sample
This is a sample gorma project. The purpose of the repository is to help you better understand the components of goa and gorma. It was created from the [gorma-design-base](https://github.com/cboecking/goa-design-base). with the following properties:

* MOEBOE_PROP_YOUR_REPO_NAME=github.com
* MOEBOE_PROP_YOUR_GIT_USER_NAME=cboecking
* MOEBOE_PROP_GOA_PACKAGE_NAME=gorma-sample
* MOEBOE_PROP_GOA_API_NAME=cellar-aaa
* MOEBOE_PROP_GOA_API_PORT=8081
* MOEBOE_PROP_GOA_RESOURCE_NAME=accountbbb
* MOEBOE_PROP_GOA_MEDIA_TYPE_NAME=Accountccc
* MOEBOE_PROP_GORMA_STORAGE_GROUP=cellar-stor
* MOEBOE_PROP_GORMA_STORAGE_MODEL=Accountddd

Notice that no two parameters are the same. This helps you understand the code better by isolating each component with a different name. Learning the framework is harder when everythig (resource, mediatype, model, etc..) is named 'account'.

Note: assumes you have postgresql installed either locally or via docker. (ACTION: show links to both)

To launch the service:

* `cd $GOPATH/src/github.com/cboecking/gorma-sample/`
* `go build .`
* `./gorma-sample`

To interact with the service:

* List accounts: `http http://localhost:8081/cellar-aaa/accountbbbs`
* Create an account: `http POST http://localhost:8081/cellar-aaa/accountbbbs name=chuck`
* List accounts: `http http://localhost:8081/cellar-aaa/accountbbbs`
* Create another account: `http POST http://localhost:8081/cellar-aaa/accountbbbs name=tom`
* Show single account: `http http://localhost:8081/cellar-aaa/accountbbbs/2`
* List accounts: `http http://localhost:8081/cellar-aaa/accountbbbs`
* Delete an account: `http DELETE http://localhost:8081/cellar-aaa/accountbbbs/2`
* Update an account: `http PUT http://localhost:8081/cellar-aaa/accountbbbs/1 name=1updated`
* List accounts: `http http://localhost:8081/cellar-aaa/accountbbbs`
