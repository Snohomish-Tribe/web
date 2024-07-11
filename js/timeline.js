/* ================================= 
    Selected Elements
==================================== */
const btn = document.querySelector(".btn");
const aside = document.querySelector("aside"); 
const col = document.querySelector(".row-3 .col-2")

let hideTimeline = true;

/* ================================= 
    Event Listeners
==================================== */
btn.addEventListener("click", () => {
    hideTimeline = !hideTimeline;    
    if (hideTimeline === false) btn.textContent = "Collapse"; 
    if (hideTimeline === true) btn.textContent = "Expand";

    col.classList.toggle("expand");
});