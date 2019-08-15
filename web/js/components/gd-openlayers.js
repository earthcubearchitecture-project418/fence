import Map from './ol/Map.js';
import View from './ol/View.js';
//import proj4 from 'proj4';
import {register} from './ol/proj/proj4.js';
import {get as getProjection} from './ol/proj.js';


(function () {
    class GDOLMap extends HTMLElement {
        constructor() {
            super();

            let shadow = this.attachShadow({mode: 'open'});
            var mapdiv = document.createElement('div');
            mapdiv.setAttribute('id', 'mapdiv');
            mapdiv.style.height = "380px";
            // mapdiv.innerHTML = ' <link rel="stylesheet" href="https://unpkg.com/leaflet@1.5.1/dist/leaflet.css"/> <script src="https://unpkg.com/leaflet@1.5.1/dist/leaflet.js" integrity="sha512-GffPMF3RvMeYyc1LWMHtK8EbPv0iNZ8/oTtHPx9/cc2ILxQ+u905qIwdpULaqDkyBKgOaB57QTMg7ztg8Jm2Og==" crossorigin=""></script> ';
            shadow.appendChild(mapdiv);

            var long = -0.09;
            var lat = 51.505;

                map = new OpenLayers.Map("mapdiv");
                map.addLayer(new OpenLayers.Layer.OSM());

                var lonLat = new OpenLayers.LonLat( -0.1279688 ,51.5077286  )
                .transform(
                        new OpenLayers.Projection("EPSG:4326"), // transform from WGS 1984
                            map.getProjectionObject() // to Spherical Mercator Projection

                );

                var zoom=16;

                var markers = new OpenLayers.Layer.Markers( "Markers"  );
                map.addLayer(markers);

                markers.addMarker(new OpenLayers.Marker(lonLat));

                map.setCenter (lonLat, zoom);


       }
    }
    window.customElements.define('geodex-olmap', GDOLMap);
})();
