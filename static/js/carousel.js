document.addEventListener("DOMContentLoaded", () => {
/* ================================= 
    Selected Elements
==================================== */
const nextBtn = document.querySelector(".next-btn");
const prevBtn = document.querySelector(".prev-btn");
const images = [...document.querySelectorAll(".carousel-img")];
const frames = [...document.querySelectorAll(".frame")];
const breadcrumbs = document.querySelector('.breadcrumbs');

/* ======= Create photo gallery navigation dots ========= */
for(let i = 0; i < images.length; i++){
  const crumb = document.createElement('div'); 
  if (i === 0) {
    crumb.classList.add('active-crumb');
  }
  crumb.classList.add('crumb');
  breadcrumbs.appendChild(crumb); 
}
const crumbs = document.querySelectorAll ('.crumb');

let translateNum = 0;
//const limit is to prevent excess white space from showing when images are moving
const limit = - (images.length - 1);

const setActiveImage = (translateNum) => {
  const index = -1 * translateNum;
  frames.forEach(
    (frame) => {
        (frame.style.transform = `translate(${translateNum * 100}%)`)
    }
  );

  //sets active navigation dot
  crumbs.forEach((crumb) => {
    crumb.classList.remove("active-crumb"); 
  }); 
  crumbs[index].classList.add("active-crumb")
}

/* ================================= 
    Event Handlers
==================================== */
nextBtn.addEventListener("click", () => {
  if(translateNum === limit) return;
  translateNum--;
  setActiveImage(translateNum); 
});

prevBtn.addEventListener("click", () => {
  if(translateNum === 0) return;  
  translateNum++;
  setActiveImage(translateNum); 
});
})