(function () {
    class FileExt extends HTMLElement {
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
            if (obj == null) {
                this.shadowRoot.innerHTML = `
                <div>
                    <span>No data graph found that we can parse</span>
                </div>
                  `;
            } else {
                var ext = obj.name.split('.').pop();

                this.shadowRoot.innerHTML = `
                    <style>
                          p {
                          }
                    </style>
                    <div>
                       <a target="_blank" href="http://geodex.org">

<svg xmlns:svg="http://www.w3.org/2000/svg" xmlns="http://www.w3.org/2000/svg" version="1.1" width="100%" height="100%" viewBox="0 0 43 46">
<defs>
  <mask id="text-clip">
    <rect id="bg" width="100%" height="100%" fill="white"/>
    <text xml:space="preserve" style="font-size:12;font-style:normal;font-family:Oswald;text-anchor:middle;text-transform:uppercase"
    x="17.992004" y="34.21402" id="text-clip-value">${ext}</text>
  </mask>
</defs>

<g id="g3308">
<path clip-rule="evenodd" d="M 36 23.29 c 0 0 0 11.43 0 11.43 c 0 1.26 -1.01 2.28 -2.25 2.28 c 0 0 -31.5 0 -31.5 0 c -1.24 0 -2.25 -1.02 -2.25 -2.28 c 0 0 0 -11.43 0 -11.43 c 0 -1.26 1.01 -2.29 2.25 -2.29 c 0 0 31.5 0 31.5 0 c 1.24 0 2.25 1.03 2.25 2.29 c 0 0 0 0 0 0 m -6.12 -18.5 c 0 0 0 8.37 0 8.37 c 0 0.53 0.43 0.96 0.96 0.96 c 0 0 8.6 0 8.6 0 c 0 0 -9.56 -9.33 -9.56 -9.33 m 13.12 12.21 c 0 0 0 25 0 25 c 0 2.21 -1.79 4 -4 4 c 0 0 -31 0 -31 0 c -2.21 0 -4 -1.79 -4 -4 c 0 0 0 -2 0 -2 c 0 0 3 0 3 0 c 0 0 0 2 0 2 c 0 0.56 0.45 1 1 1 c 0 0 31 0 31 0 c 0.55 0 1 -0.44 1 -1 c 0 0 0 -25 0 -25 c 0 0 -9.16 0 -9.16 0 c -2.12 0 -3.84 -1.71 -3.84 -3.84 c 0 0 0 -10.14 0 -10.14 c -0.06 0 -0.12 -0.02 -0.19 -0.02 c 0 0 -18.81 0 -18.81 0 c -0.55 0 -1 0.45 -1 1 c 0 0 0 14 0 14 c 0 0 -3 0 -3 0 c 0 0 0 -14 0 -14 c 0 -2.21 1.79 -4 4 -4 c 0 0 18.81 0 18.81 0 c 2.63 0 4 1.5 4 1.5 c 0 0 11.07 10.81 11.07 10.81 c 0 0 0.57 0.5 0.89 1.81 c 0 0 0.23 0 0.23 0 c 0 0 0 2.88 0 2.88 c 0 0 0 0 0 0" id="path313" mask="url(#text-clip)"/>

</g>
</svg>

                </div>
                      `;
            }
        }
    }
    window.customElements.define('geodex-fileext', FileExt);
})();
