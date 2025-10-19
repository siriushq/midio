const current = document.currentScript;
if (!current || current.type === "module") throw new Error();

fetch("https://midio.sirius.cam/theme/index.mjs").then(contents => {
    const element = document.createElement("script");
    element.type = "module";
    element.innerHTML = contents;
    current.insertAdjacentElement("afterend", element);
})