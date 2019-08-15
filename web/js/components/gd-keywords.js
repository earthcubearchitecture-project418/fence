(function () {
    class GeoKeywords extends HTMLElement {
        constructor() {
            super();

            var printError = function (error, explicit) {
                console.log(`[${explicit ? 'EXPLICIT' : 'INEXPLICIT'}] ${error.name}: ${error.message}`);
            };

            // need to think about calling jsonld.js and using
            // it to parse the graph
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

            this.attachShadow({ mode: 'open' });
            if (obj ==  null) {
                this.shadowRoot.innerHTML = `
                <div>
                    <span>No data graph found that we can parse</span>
                </div>
                  `;
            } else {
                this.shadowRoot.innerHTML = `
                <div>
                    ${obj.keywords}
                </div>
                  `;
            }

        }
    }
    window.customElements.define('geodex-keywords', GeoKeywords);
})();