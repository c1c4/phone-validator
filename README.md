# Phone Validator
Web App to show you some phones and show if it's valid or not.

<img width="1679" alt="Screen Shot 2022-02-04 at 17 48 39" src="https://user-images.githubusercontent.com/5156931/152605691-7cb98d41-b9ed-4a2b-8da4-ef54e0c7ff56.png">


#
## The Api

The api is built with: 

### API
1.  [*Golang*](https://go.dev/)
1.  [*Gin*](https://github.com/gin-gonic/gin)
1.  [*Godotenv*](https://github.com/joho/godotenv)

### DB
1.  [*Gorm*](https://gorm.io/index.html)

### Test
1.  [*Testify*](https://github.com/stretchr/testify)


#

## Running the API
You can run the phone validator api with two ways:

1.  Docker for this you will need docker installed in your machine [Docker](https://www.docker.com/)

        make build.application
    
    or if you already build

        make start.application

    This will put the api online.
    
    You can use:


        docker ps

    You should see an output that starts with something that looks like the following:
    
    CONTAINER ID | IMAGE
    ------------ | -----
    a123bc007edf | api_app


    with this and I believe you should be fine and start to using the api in this url **localhost:8080**

1.  Open the project in you preferred IDE Goland or VSCode then run in your terminal:
    
        go mod download

    The database used on this project is on api folder.

#

## Test the API
Well you can test the phone validator api with two ways as well no surprises I hope:

1.  These are the commands you need and you see above
    
    For unit tests:

        make test.unit
    
    For integration tests:
        
        make test.integration

1.  You can open the project in Goland or VSCode or the IDE your like but I know these two has support for test and configure their tests
    1. [Golang on Goland](https://www.jetbrains.com/go/)
    1. [Golang on VSCode](https://code.visualstudio.com/docs/languages/go)

    Or you can open you terminal go to the project folder and run

        go test ./... -v <- this will run all tests

    You only need to make sure if you install the all the dependencies and has your .env setup.
    
#
## The Web
The api is built with: 

### Web
1.  [*Reactjs*](https://reactjs.org/)
1.  [*Yarn*](https://yarnpkg.com/)
1.  [*Webpack*](https://webpack.js.org/)
1.  [*Babel*](https://babeljs.io/)
1.  [*Prettier*](https://prettier.io/)
1.  [*Typescript*](https://www.typescriptlang.org/)


#

## Running the Web
First you need to install all the dependencies so in the terminal you can run
    
    yarn install
    
after that you can run:
    
    yarn start
    
#### Just remember to start the backend first.
