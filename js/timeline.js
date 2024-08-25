/* ================================= 
    Selected Elements
==================================== */
const btn = document.querySelector(".btn");
const aside = document.querySelector("aside"); 
const col = document.querySelector(".row-3 .col-2");

let hideTimeline = true;

/* ================================= 
    Event Listeners
==================================== */
/*  
btn.addEventListener("click", () => {
    hideTimeline = !hideTimeline;

    if (hideTimeline === false) {
        btn.textContent = 'Collapse';
        col.style.height = `${aside.offsetHeight}px`;
        col.classList.add('expand')
    } 
    
    if (hideTimeline === true) {
        btn.textContent = 'Expand';
        col.classList.remove('expand')
        col.style.height = '450px'
    }
});*/

btn.addEventListener("click", () => {
    hideTimeline = !hideTimeline;

    if (hideTimeline === false) {
        btn.textContent = 'Collapse';
        col.classList.add('expand')
    } 
    
    if (hideTimeline === true) {
        btn.textContent = 'Expand';
        col.classList.remove('expand')
        col.classList.remove('full')
    }
});

col.addEventListener('transitionend', () => {
    if(col.classList.contains('expand')) {
        col.classList.add('full')
    }
})


// window.addEventListener('resize', () => {
//     if(col.classList.contains('expand')) {
//         col.style.height = `${aside.offsetHeight}px`;
//     }
// })