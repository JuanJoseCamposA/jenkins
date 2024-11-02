pipeline {
    agent any
    stages {
        stage('Preparación') {
            steps {
                script {
                    // Detener y eliminar contenedores antiguos
                    sh 'docker stop charming_knuth || true'
                    sh 'docker rm charming_knuth || true'
                    sh 'rm -rf temp_jenkins || true'
                }
            }
        }
        stage('Clonar Repositorio') {
            steps {
                script {
                    // Clonar el repositorio
                    sh 'git clone -b main https://github.com/JuanJoseCamposA/jenkins.git temp_jenkins'
                }
            }
        }
        stage('Construir') {
            steps {
                dir('temp_jenkins') {
                    // Construir la imagen Docker
                    sh 'docker build -t juan ./'
                }
            }
        }
        stage('Ejecutar API en Contenedor') {
            steps {
                script {
                    // Iniciar el contenedor
                    sh 'docker run -d --name charming_knuth -p 8081:8081 juan'
                }
            }
        }
stage('Probar API') {
    steps {
        script {
            // Esperar un momento para que el contenedor se inicie
            sh 'sleep 20'
            // Probar la API
            sh 'docker logs charming_knuth'  // Ver logs del contenedor
            def responseCode = sh(script: 'curl -s -o /dev/null -w "%{http_code}" http://172.19.146.241:8081/actividades', returnStdout: true).trim()
            echo "Código de respuesta: ${responseCode}"
        }
    }
}

        stage('Desplegar') {
            steps {
                echo 'Desplegando...'
            }
        }
    }
    post {
        failure {
            echo 'La construcción falló.'
        }
    }
}
