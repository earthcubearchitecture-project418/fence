import {
	html,
	render
} from './lit-html.js';

(function () {
    class Tangram extends HTMLElement {
        constructor() {
            super();

            // ref: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
            // https://scotch.io/tutorials/how-to-use-the-javascript-fetch-api-to-get-data
            // https://stackoverflow.com/questions/45018338/javascript-fetch-api-how-to-save-output-to-variable-as-an-object-not-the-prom
            // https://stackoverflow.com/questions/38869197/fetch-set-variable-with-fetch-response-and-return-from-function

            var printError = function (error, explicit) {
                console.log(`[${explicit ? 'EXPLICIT' : 'INEXPLICIT'}] ${error.name}: ${error.message}`);
            };

            // GET test
            function tj_providers(id) {
                return fetch(id)
                    .then(function (response) {
                        return response.text();
                    })
                    .then(function (bodytext) {
                        //  console.log(id);
                        // console.log(bodytext);
                        // return JSON.stringify(bodytext);
                        return bodytext;
                    });
            }

            // TODO read in a component attribute to get the URL passed
            // can a component read the URL of the parent page?

            const shape = this.getAttribute('google-shape');
            console.log(shape);
            var url = new URL(window.location.href);
            var urlparam = url.searchParams.get("url");
            var geturl = "https://tangram.gleaner.io/ucheck?url=" +  urlparam + "&format=human&shape=" + shape;

            // GET test call...
            tj_providers(geturl).then((providers) => {
                // do something with the providers
                this.attachShadow({ mode: 'open' });
                // var h =  `<div>${itemTemplates}</div>`;
                this.shadowRoot.innerHTML = `<div style="width=100%;overflow:auto;"><pre>${providers}</pre></div>` ;
                // this.shadowRoot.appendChild(this.cloneNode(h));
            });
        }
    }
    window.customElements.define('gleaner-tangram', Tangram);
})();


