Jenkinsfile----CODIGO DE LA PRUEBA


![image](https://github.com/user-attachments/assets/c21abdb3-2dd1-4fea-bd0f-edff3173a639)

![image](https://github.com/user-attachments/assets/86e1329e-7247-4c0e-b69e-0f55d2a4b241)

![image](https://github.com/user-attachments/assets/047fc317-3c31-4f80-a563-14096ad18af1)

![image](https://github.com/user-attachments/assets/5e0f6fa5-bca0-40f0-ad1b-babf5ed4c2f6)
![image](https://github.com/user-attachments/assets/458202d6-d77b-4d90-a6e5-4b0e283cf229)
![image](https://github.com/user-attachments/assets/97094c03-47ce-4999-a098-3078e81385d1)
![image](https://github.com/user-attachments/assets/f05fe95a-0279-495c-91dd-0889d47e8d5f)
![image](https://github.com/user-attachments/assets/19b00dfe-c36b-44d5-a6de-bf653076bcd0)
![image](https://github.com/user-attachments/assets/371652da-1f50-4487-a6f3-6cbb13857295)
![image](https://github.com/user-attachments/assets/e13153ec-0c68-4b00-8441-1b4c79d00ba1)
![image](https://github.com/user-attachments/assets/5dc7ed8a-dd59-4853-9ba2-0cb6cc04741c)
![image](https://github.com/user-attachments/assets/c28fecd9-a7d7-49fa-abb1-612db34d0481)
![image](https://github.com/user-attachments/assets/ffa9f738-ee14-4478-97bb-d33a177d3f03)


Este pipeline de Jenkins implementa un proceso de **Integración Continua/Despliegue Continuo (CI/CD)** para una API. A continuación se describen las etapas del pipeline:

1. Preparación:
   - Detiene y elimina cualquier contenedor Docker que esté utilizando el nombre `charming_knuth`.
   - Elimina la carpeta `temp_jenkins` si existe, para asegurarse de que el entorno esté limpio antes de comenzar.

2. Clonar Repositorio:
   - Clona el repositorio de GitHub desde la rama `main` en un directorio temporal llamado `temp_jenkins`.

3. Construir:
   - Cambia al directorio `temp_jenkins` y construye una imagen Docker utilizando el Dockerfile presente en ese directorio. La imagen se etiqueta como `juan`.

4. Ejecutar API en Contenedor:
   - Inicia un contenedor en segundo plano a partir de la imagen `juan`, asignando el nombre `charming_knuth` y mapeando el puerto `8081` del contenedor al puerto `8081` de la máquina host.

5. Probar API:
   - Espera 20 segundos para asegurarse de que el contenedor esté completamente iniciado.
   - Verifica los logs del contenedor para confirmar que el servidor de la API se está ejecutando.
   - Realiza una solicitud `curl` al endpoint `/actividades` de la API utilizando la dirección IP del host (en este caso, `http://172.19.146.241:8081/actividades`) para comprobar su disponibilidad. Si la respuesta es un código `200`, significa que la API está funcionando correctamente.

6. Desplegar:
   - Esta etapa simplemente imprime un mensaje indicando que el despliegue ha ocurrido.



