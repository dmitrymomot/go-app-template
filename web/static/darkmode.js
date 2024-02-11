// Function to add a class at the beginning of the class list of the documentElement
function addClassAtStart(className) {
  var currentClassValue = document.documentElement.className;

  // Check if the class already exists in the class list to avoid duplicates
  var classList = currentClassValue.split(' ');
  if (classList.indexOf(className) === -1) {
    // Add the new class at the start of the current class value
    document.documentElement.className = className + ' ' + currentClassValue;
  }
}


// On page load or when changing themes, best to add inline in `head` to avoid FOUC
if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  addClassAtStart('dark')
} else {
  document.documentElement.classList.remove('dark')
}

// // Whenever the user explicitly chooses light mode
// localStorage.theme = 'light'

// // Whenever the user explicitly chooses dark mode
// localStorage.theme = 'dark'

// // Whenever the user explicitly chooses to respect the OS preference
// localStorage.removeItem('theme')
