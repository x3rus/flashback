
pipeline {
    agent any

        stages {

            stage('Unittest') {
                steps {
                    /* use unittest */
                    script {
                       sh 'make test'
                    }
                }
            }

            stage('Build') {
                steps {
                    script {
                     sh 'make build'   
                    }
                }
            }
            stage('Lint') {
                steps {
                    script {
                     sh 'make lint'   
                    }
                }
            }


        } // END stages

    post {
        success {
            script {
                echo 'slack me '
            }
        }
        failure {
            echo 'slack me with a failure message'
        }
    } // END POST       

} // END pipeline 
