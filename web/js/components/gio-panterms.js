/*jshint esversion: 6 */
import {
    html,
    render
} from './lit-html.js';

(function () {
    class Panterms extends HTMLElement {
        constructor() {
            super();

            // Error setup
            var printError = function (error, explicit) {
               console.log(`[${explicit ? 'EXPLICIT' : 'INEXPLICIT'}] ${error.name}: ${error.message}`);
            };

            function tj_providers(res) {
                return fetch(`${res}`, {
                       headers: {
                         'Content-Type': 'application/json',
                         'Accept': 'application/json'
                       }
                }).then((response) => response.json());
            }

            // let promise = new Promise(function () {
            // async function getUserAsync() {
            // Get the keywords from the data graph
            var obj;
            var inputs = document.getElementsByTagName('script');
            for (var i = 0; i < inputs.length; i++) {
                if (inputs[i].type.toLowerCase() == 'application/ld+json') {
                    try {
                        obj = JSON.parse(inputs[i].innerHTML);
                    } catch (e) {
                        if (e instanceof SyntaxError) {
                            printError(e, true);
                        } else {
                            printError(e, false);
                        }
                    }
                }
            }

            var terms = obj.keywords.split(" ");
            const itemTemplates = [];
             // GET test call...
             this.attachShadow({ mode: 'open' });

            for (var j = 0; j < terms.length; j++) {
                var geturl = "http://seprojects.marum.de:8383/pangterm/api?name=" + terms[j];
                tj_providers(geturl).then(providers => {
                    console.log(providers.parameter);
                    // console.log();
                    let du = providers.text_match[0].term[0].description_uri;
                    let score = providers.text_match[0].term[0].score;
                    console.log(du);
                    itemTemplates.push(html`<div>${providers.parameter} (score: ${score} )  : <a target="_blank" href="${du}">${du}<a></div>`);
                    var h = html`<div style="margin-top:10px">${itemTemplates}</div>`;        
                    render(h, this.shadowRoot);
                });
            }

        }
    }
    window.customElements.define('gio-panterms', Panterms);
})();


