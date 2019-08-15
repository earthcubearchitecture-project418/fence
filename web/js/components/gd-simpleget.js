import {
	html,
	render
} from './lit-html.js';

(function () {
    class SimpleGet extends HTMLElement {
        constructor() {
            super();

            // ref: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
            // https://scotch.io/tutorials/how-to-use-the-javascript-fetch-api-to-get-data
            // https://stackoverflow.com/questions/45018338/javascript-fetch-api-how-to-save-output-to-variable-as-an-object-not-the-prom
            // https://stackoverflow.com/questions/38869197/fetch-set-variable-with-fetch-response-and-return-from-function

            // GET test
            function tj_providers(id) {
                return fetch('http://geodex.org/api/v1/typeahead/providers')
                    .then(function (response) {
                        return response.json();
                    })
                    .then(function (myJson) {
                         console.log(id);
                        // console.log(JSON.stringify(myJson));
                        // return JSON.stringify(myJson);
                        return myJson;
                    });
            }

            // GET test call...
            tj_providers("We can pass values..  ").then((providers) => {
                // do something with the providers
                this.attachShadow({ mode: 'open' });

                var count = Object.keys(providers).length;
                const itemTemplates = [];
                var i;
                for (i = 0; i < count; i++) {
                    // console.log(providers[i].name)
                    itemTemplates.push(  `${providers[i].name}`);
                    // console.log(itemTemplates)
                }

                var h =  `<div>${itemTemplates}</div>`;
                this.shadowRoot.innerHTML = `${h}` ;
                // this.shadowRoot.appendChild(this.cloneNode(h));
            });

        }
    }
    window.customElements.define('geodex-simpleget', SimpleGet);
})();


