const script = `
import SimpleBar from "https://esm.run/simplebar";
import "https://esm.run/simplebar/dist/simplebar.css";

import ResizeObserver from "https://esm.run/resize-observer-polyfill";
window.ResizeObserver = ResizeObserver;


const sidebar = document.getElementById("sidebar");
new SimpleBar(sidebar);

const noise = document.createElement("div");
noise.className = "background-noise";

const overlay = document.createElement("div");
overlay.className = "background-filter";

const image = document.createElement("img");
image.className = "background";
image.src = "https://midio.sirius.cam/theme/background.webp";
image.draggable = false;

const marker = document.querySelector('a[href="#content-area"]');
if (!marker) throw new Error("could not apply background style");
marker.insertAdjacentElement("afterend", image);
marker.insertAdjacentElement("afterend", overlay);
marker.insertAdjacentElement("afterend", noise);
`;

const current = document.currentScript;
if (!current || current.type === "module") throw new Error();

const element = document.createElement("script");
element.type = "module";
element.innerHTML = script;
current.insertAdjacentElement("afterend", element);