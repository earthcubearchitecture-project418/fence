(function () {
    class ReqShape extends HTMLElement {
        constructor() {
            super();

            // ref: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
            // https://scotch.io/tutorials/how-to-use-the-javascript-fetch-api-to-get-data
            // https://stackoverflow.com/questions/45018338/javascript-fetch-api-how-to-save-output-to-variable-as-an-object-not-the-prom
            // https://stackoverflow.com/questions/38869197/fetch-set-variable-with-fetch-response-and-return-from-function

            // POST test call
            async function tj_postData(url = ``, data = {}) {
                // Default options are marked with *
                return fetch(url, {
                    method: "POST", // *GET, POST, PUT, DELETE, etc.
                    mode: "no-cors", // no-cors, cors, *same-origin
                    cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
                    credentials: "same-origin", // include, *same-origin, omit
                    headers: {
                        "Content-Type": "application/json", // "Content-Type": "application/x-www-form-urlencoded",
                    },
                    redirect: "follow", // manual, *follow, error
                    referrer: "no-referrer", // no-referrer, *client
                    body: JSON.stringify(data), // body data type must match "Content-Type" header
                })
                .then(function (response) {
                    console.log(response.json);
                    return response.json();
                })
                .then(function (myjson) {
                    console.log(myjson);
                    return JSON.stringify(myjson);
                });
                // .then(response => response.json()); // parses JSON response into native Javascript objects 
            }


            // try https://alligator.io/js/fetch-api/


            // POST implementation call
            tj_postData(`https://jsonplaceholder.typicode.com/posts`, { title: 'foo', body: 'bar', userId: 1 }).then((r) => {
                // do something with the providers
                this.attachShadow({ mode: 'open' });
                this.shadowRoot.innerHTML = `
            <div>
              POST  Results:  ${JSON.stringify(r)}
            </div>
              `;
            });

        }
    }
    window.customElements.define('geodex-shaclreq', ReqShape);
})();


