let delay = [0, 400]

function handleLoader(i) {
  if (i < 2) {
    setTimeout(function() {
      document.querySelector(".page-load").classList.add("pl-" + i)
      handleLoader(i + 1)
    }, delay[i])
  }
}

const hideLoader = () => handleLoader(0)

export default hideLoader
