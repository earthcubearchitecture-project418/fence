(function () {
    class GeoCitation extends HTMLElement {
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

            var today = new Date();
            var dd = String(today.getDate()).padStart(2, '0');
            var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
            var yyyy = today.getFullYear();

            today = mm + '/' + dd + '/' + yyyy;

            var version = 'Not Provided'; // override if set by the SDO

            //  still need  <span> Distribution org, </span>  <span> Release Date, </span>
            this.attachShadow({ mode: 'open' });
            if (obj == null) {
                this.shadowRoot.innerHTML = `
                <div>
                    <span>No data graph found that we can parse</span>
                </div>
                  `;
            } else {
                this.shadowRoot.innerHTML = `
                    <div style="overflow-wrap: break-word;width=100%">
                         ${obj.publisher.name},
                          ${obj.name},   
                         Version is ${version}  
                         Data set accessed ${today}  
                         at <a href='${obj.distribution.contentUrl}'>${obj.distribution.contentUrl}</a>, 
                    </div>
                      `;
            }
        }
    }
    window.customElements.define('geodex-citation', GeoCitation);
})();
