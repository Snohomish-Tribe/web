/* ================================= 
    Selected Elements
==================================== */
const nextBtn = document.querySelector(".next-btn");
const prevBtn = document.querySelector(".prev-btn");
const images = [...document.querySelectorAll(".carousel-img")];
const crumbs = [...document.querySelectorAll(".crumb")];

let translateNum = 0;

const setActiveImage = (translateNum) => {
  const index = -1 * translateNum;
  images.forEach(
    (img) => {
        (img.style.transform = `translate(${translateNum * 100}%)`)
    }
  );

  crumbs.forEach((crumb) => {
    crumb.classList.remove("active-crumb"); 
  }); 
  
  crumbs[index].classList.add("active-crumb")
}

/* ================================= 
    Event Handlers
==================================== */
nextBtn.addEventListener("click", () => {
  if(translateNum === -3) return;
  translateNum--;
  setActiveImage(translateNum); 
});

prevBtn.addEventListener("click", () => {
  if(translateNum === 0) return;  
  translateNum++;
  setActiveImage(translateNum); 
});
