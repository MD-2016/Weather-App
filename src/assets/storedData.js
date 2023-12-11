// adding code to store the user's session
const pageReloadCheck = (
    (window.performance
        .getEntriesByType('navigation')
        .map((nav) => nav.type)
        .includes('reload')    
    )
);

alert(pageReloadCheck)