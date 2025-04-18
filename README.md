# quiz-service-clone

A self hostable kahoot clone written in Svelte and Go

### Built With

* [![Svelte][Svelte.dev]][Svelte-url]
* [![Go][Go.dev]][Go-url]

### Installation

1. Clone the repo and navigate to it
    ``` bash
    $ git clone https://github.com/jhands0/quiz-service-clone.git quiz-service-clone

    $ cd quiz-service-clone
    ```

2. Build the backend docker image
    ``` bash
    $ docker build --file=backend/backend.dockerfile -t quiz-clone-backend .
    ```

3. Build the frontend docker image
    ``` bash
    $ docker build --file=backend/frontend.dockerfile -t quiz-clone-frontend .
    ```

4. Run the docker-compose.yaml file
    ``` bash
    $ docker-compose -f docker-compose.yaml up
    ```



[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Go.dev]: https://img.shields.io/badge/Go-4A4A55?style=for-the-badge&logo=go&logoColor=00ADD8
[Go-url]: https://go.dev/