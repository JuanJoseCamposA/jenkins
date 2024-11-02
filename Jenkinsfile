pipeline {
    agent any  // Usar cualquier agente disponible para ejecutar el pipeline
    
    stages {
        stage('Preparación') {  // Etapa de preparación
            steps {
                script {
                    // Detener y eliminar contenedores antiguos para evitar conflictos
                    sh 'docker stop charming_knuth || true'  // Detiene el contenedor si está en ejecución
                    sh 'docker rm charming_knuth || true'  // Elimina el contenedor
                    sh 'rm -rf temp_jenkins || true'  // Elimina la carpeta temporal de Jenkins
                }
            }
        }
        
        stage('Clonar Repositorio') {  // Etapa para clonar el repositorio de GitHub
            steps {
                script {
                    // Clonar el repositorio en la carpeta temporal
                    sh 'git clone -b main https://github.com/JuanJoseCamposA/jenkins.git temp_jenkins'
                }
            }
        }
        
        stage('Construir') {  // Etapa para construir la imagen Docker
            steps {
                dir('temp_jenkins') {  // Cambia al directorio del repositorio clonado
                    // Construir la imagen Docker usando el Dockerfile presente en el directorio
                    sh 'docker build -t juan ./'
                }
            }
        }
        
        stage('Ejecutar API en Contenedor') {  // Etapa para ejecutar la API en un contenedor Docker
            steps {
                script {
                    // Iniciar el contenedor en segundo plano, mapeando el puerto 8081
                    sh 'docker run -d --name charming_knuth -p 8081:8081 juan'
                }
            }
        }

        stage('Probar API') {  // Etapa para probar la API después de que el contenedor está en funcionamiento
            steps {
                script {
                    // Esperar un momento para permitir que el contenedor se inicie correctamente
                    sh 'sleep 20'
                    // Ver los logs del contenedor para verificar que se esté ejecutando correctamente
                    sh 'docker logs charming_knuth'  
                    // Probar la API utilizando cURL y capturar el código de respuesta HTTP
                    def responseCode = sh(script: 'curl -s -o /dev/null -w "%{http_code}" http://172.19.146.241:8081/actividades', returnStdout: true).trim()
                    // Imprimir el código de respuesta en la consola
                    echo "Código de respuesta: ${responseCode}"
                }
            }
        }

        stage('Desplegar') {  // Etapa para realizar despliegues adicionales si es necesario
            steps {
                echo 'Desplegando...'  // Mensaje de despliegue, se puede agregar lógica de despliegue aquí
            }
        }
    }
    
    post {
        failure {  // Acción que se ejecuta si la construcción falla
            echo 'La construcción falló.'  // Mensaje de error en la consola
        }
    }
}
