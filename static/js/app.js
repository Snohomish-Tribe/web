/* =================================
    Selected Elements
==================================== */
const menuBtn = document.querySelector(".toggle");
const mainNav = document.querySelector(".main-nav");
const menuIcn = document.querySelector(".toggle img");

let isMenuOpen = false;

/* =================================
    Event Listeners
==================================== */
menuBtn.addEventListener("click", (event) => {
  isMenuOpen = !isMenuOpen;
  const btn = event.currentTarget;
  if (btn.classList.contains("toggle")) {
    mainNav.classList.toggle("show-links");
  }

  if (isMenuOpen === true) menuIcn.src = "images/close.svg"; 
  if (isMenuOpen === false) menuIcn.src = "images/hamburger.svg"; 
  
});

console.log("app")