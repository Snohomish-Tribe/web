/* =================================
    Selected Elements
==================================== */
const menuBtn = document.querySelector(".toggle");
const mainNav = document.querySelector(".main-nav");
const closeIcn = document.querySelector(".close");
const menuIcn = document.querySelector(".menu");

let isMenuOpen = false;

/* =================================
    Event Listeners
==================================== */
menuBtn.addEventListener("click", (event) => {
  isMenuOpen = !isMenuOpen;
  const btn = event.currentTarget;
  if (btn.classList.contains("toggle")) {
    mainNav.classList.toggle("show-links");
    closeIcn.classList.toggle("hide");
    menuIcn.classList.toggle("hide");
  }
});
