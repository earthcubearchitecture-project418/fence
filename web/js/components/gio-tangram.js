/*jshint esversion: 6 */
import {
	html,
	render
} from './lit-html.js';

(function () {
    class Tangram extends HTMLElement {
        constructor() {
            super();

            // Error setup
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

            const shape = this.getAttribute('google-shape');
            console.log(shape);
            var url = new URL(window.location.href);
            var urlparam = url.searchParams.get("url");

            if (urlparam) {
                var geturl = "https://tangram.gleaner.io/ucheck?url=" +  urlparam + "&format=human&shape=" + shape;

                // GET test call...
                tj_providers(geturl).then((providers) => {
                    // do something with the providers
                    this.attachShadow({ mode: 'open' });
                    // var h =  `<div>${itemTemplates}</div>`;
                    // this.shadowRoot.appendChild(this.cloneNode(h));
                    if (providers.indexOf("Conforms: True") !== -1) {
                        this.shadowRoot.innerHTML = `<div style="background:#F8FFD4;width=100%;overflow:auto;"><pre>${providers}</pre></div>` ;
                    }
                    if (providers.indexOf("Conforms: False") !== -1) {
                        this.shadowRoot.innerHTML = `<div style="background:#FFDCD4;width=100%;overflow:auto;"><pre>${providers}</pre></div>` ;
                    }
                });
            }
            else {
                this.attachShadow({ mode: 'open' });
                // var h =  `<div>${itemTemplates}</div>`;
                this.shadowRoot.innerHTML = `<div style="width=100%;overflow:auto;"><pre>No URL provided to check</pre></div>` ;
            }


        }
    }
    window.customElements.define('gleaner-tangram', Tangram);
})();


