/* =================================
    Selected Elements
==================================== */
const menuBtn = document.querySelector(".toggle");
const mainNav = document.querySelector(".main-nav");
const closeIcn = document.querySelector(".close");
const menuIcn = document.querySelector(".menu");
const year = document.querySelector('.year');

let isMenuOpen = false;

//Displays year in the footer
year.textContent = year.textContent = `@ ${new Date().getFullYear()} Snohomish Tribe of Indians`;

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
