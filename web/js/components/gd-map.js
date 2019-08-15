import * as L from './leaflet/leaflet-src.esm.js';

(function () {
    class GDMap extends HTMLElement {
        constructor() {
            super();

            let shadow = this.attachShadow({mode: 'open'});
            var mapdiv = document.createElement('div');
            mapdiv.setAttribute('id', 'mapid');
            // mapdiv.setAttribute('style', 'height:180px');
            mapdiv.style.height = "180px";
            //mapdiv.innerHTML = ' <link rel="stylesheet" href="js/components/leaflet/leaflet.css" ';
            //mapdiv.innerHTML = ' <link rel="stylesheet" href="https://unpkg.com/leaflet@1.5.1/dist/leaflet.css" /> ';
            mapdiv.innerHTML = ' <link rel="stylesheet" href="https://unpkg.com/leaflet@1.5.1/dist/leaflet.css"/> <script src="https://unpkg.com/leaflet@1.5.1/dist/leaflet.js" integrity="sha512-GffPMF3RvMeYyc1LWMHtK8EbPv0iNZ8/oTtHPx9/cc2ILxQ+u905qIwdpULaqDkyBKgOaB57QTMg7ztg8Jm2Og==" crossorigin=""></script> ';
            shadow.appendChild(mapdiv);

            var long = -0.09;
            var lat = 51.505;
            var map = L.map(mapdiv).setView([lat,long], 4);

            L.marker([lat,long]).addTo(map)
                .bindPopup('A pretty CSS3 popup.<br> Easily customizable.')
                .openPopup();

            L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributor(s)'
            }).addTo(map);

       }
    }
    window.customElements.define('geodex-map', GDMap);
})();
