import {
	html,
	render
} from './lit-html.js';

import Map from './ol/Map.js';
import View from './ol/View.js';
import {register} from './ol/proj/proj4.js';
import {get as getProjection} from './ol/proj.js';
import TileLayer from './ol/layer/Tile.js';
import OSM from './ol/source/OSM.js';

(function () {
    class GDOLMap extends HTMLElement {
        constructor() {
            super();

            // let shadow = this.attachShadow({mode: 'open'});
            // var mapdiv = document.createElement('div');
            // mapdiv.setAttribute('id', 'mapdiv');
            // mapdiv.style.height = "380px";
            // shadow.appendChild(mapdiv);

            this.attachShadow({mode: 'open'});
            this.shadowRoot.innerHTML = `<div id="map"></div>`;

            var long = -0.09;
            var lat = 51.505;

            var map = new Map({
                layers: [
                    new TileLayer({
                                    source: new OSM()

                    })

                ],
                        target: 'map',
                view: new View({
                              center: [0, 0],
                              zoom: 2

                })

            });

       }
    }
    window.customElements.define('geodex-olmap', GDOLMap);
})();
