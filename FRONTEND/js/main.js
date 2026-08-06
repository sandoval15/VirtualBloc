document.addEventListener('DOMContentLoaded', () => {
    let toggle = document.querySelector('.toggle-bookshelf')

    toggle.addEventListener('click', () => {
        let style = document.documentElement.style
        if (toggle.getAttribute('hide') === 'true') {
            style.setProperty('--bookshelf-radius', '10px')
            style.setProperty('--bookshelf-move', '300px')
            style.setProperty('--toggle-rot', '180deg')
            style.setProperty('--toggle-offset', '0px')
            toggle.setAttribute('hide', 'false')
        } else {
            style.setProperty('--bookshelf-radius', '25%')
            style.setProperty('--bookshelf-move', '0px')
            style.setProperty('--toggle-rot', '360deg')
            style.setProperty('--toggle-offset', '-10px')
            toggle.setAttribute('hide', 'true')
        }
    })
})