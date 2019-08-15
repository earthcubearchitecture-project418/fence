(function () {
    class GeoDexLogo extends HTMLElement {
        constructor() {
            super();

            // need to think about calling jsonld.js and using
            // it to parse the graph
            var obj;
            var inputs = document.getElementsByTagName('script');
            for (var i = 0; i < inputs.length; i++) {
                if (inputs[i].type.toLowerCase() == 'application/ld+json') {
                    obj = JSON.parse(inputs[i].innerHTML);
                }
            }

            this.attachShadow({ mode: 'open' });
            this.shadowRoot.innerHTML = `
                    <style>
                          p {
                          }
                    </style>
                    <div>
                       <a target="_blank" href="http://geodex.org">
                        <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="100" height="100" viewBox="0 0 100 100">
                   <g>
                      <rect width="100%" height="100%" fill="#FFFFFF" fill-opacity="0.0"/>
                      <g transform="translate(50 50) scale(0.69 0.69) rotate(0) translate(-50 -50)" style="fill:#1A1A1A">
                               <svg xmlns:cc="http://creativecommons.org/ns#" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd" xmlns:svg="http://www.w3.org/2000/svg" fill="#1A1A1A" version="1.1" x="0px" y="0px" viewBox="0 0 100 100">
                            <g transform="translate(0,-952.36218)">
                                           <path style="text-indent:0;text-transform:none;direction:ltr;block-progression:tb;baseline-shift:baseline;color:#000000;enable-background:accumulate;" d="m 49.781198,963.38237 a 1.0001,1.0001 0 0 0 -0.1562,0.0625 l -35,14.375 a 1.0001,1.0001 0 0 0 -0.625,0.9375 l 0,47.21873 a 1.0001,1.0001 0 0 0 0.625,0.9375 l 35,14.375 a 1.0001,1.0001 0 0 0 0.75,0 l 35.000003,-14.375 a 1.0001,1.0001 0 0 0 0.625,-0.9375 l 0,-47.21873 a 1.0001,1.0001 0 0 0 -0.625,-0.9375 l -35.000002,-14.375 a 1.0001,1.0001 0 0 0 -0.5938,-0.0625 z m 0.2188,2.0625 32.375003,13.3125 -32.375002,13.28125 -32.375,-13.28125 32.375,-13.3125 z m -34,14.78125 33,13.5625 0,45.09378 -33,-13.5625 0,-45.09378 z m 68.000003,0 0,45.09378 -33.000002,13.5625 0,-45.09378 33.000002,-13.5625 z" fill="#1A1A1A" fill-opacity="1" stroke="none" marker="none" visibility="visible" display="inline" overflow="visible" />
                            </g>
                            <g>
                              <text x="25" y="60" fill="green">	&#10004;</text>
                              </g>
                         </svg>
                         </a>
                      </g>
                   </g>
                </svg>
                </div>
                      `;
        }
    }
    window.customElements.define('geodex-logo', GeoDexLogo);
})();
