import {
	html,
	render
} from './lit-html.js';

(function () {
    class Panterms extends HTMLElement {
        constructor() {
            super();

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
                        return JSON.parse(bodytext);
                    });
            }

            var term = "ocean";

            if (term) {
                // http://seprojects.marum.de:8383/pangterm/api?name=Fugacity%20of%20carbon%20dioxide%20in%20seawater
                var geturl = "http://seprojects.marum.de:8383/pangterm/api?name=" + term

                // GET test call...
                tj_providers(geturl).then((providers) => {
                    // do something with the providers
                    this.attachShadow({ mode: 'open' });
                    // var h =  `<div>${itemTemplates}</div>`;
                    this.shadowRoot.innerHTML = `<div style="width=100%;overflow:auto;"><pre>${providers.parameter}</pre></div>` ;
                    // this.shadowRoot.appendChild(this.cloneNode(h));
                });
            }
            else {
                this.attachShadow({ mode: 'open' });
                // var h =  `<div>${itemTemplates}</div>`;
                this.shadowRoot.innerHTML = `<div style="width=100%;overflow:auto;"><pre>No URL provided to check</pre></div>` ;
            }


        }
    }
    window.customElements.define('gio-panterms', Panterms);
})();


