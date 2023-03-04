let barsIsOpen = false;

function openBars() {
  let hamburgerNav = document.getElementById('hamburger-nav');

  if (barsIsOpen) {
    hamburgerNav.style.display = 'none';
    barsIsOpen = false;
  } else {
    hamburgerNav.style.display = 'block';
    barsIsOpen = true;
  }
}
