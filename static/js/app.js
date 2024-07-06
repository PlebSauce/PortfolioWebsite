const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
        console.log(entry)
        if(entry.isIntersecting) {
            entry.target.classList.add('show');
        }
        else {
            entry.target.classList.remove('show');
        }
    });
});

const hiddenElements = document.querySelectorAll('.list-projects');
hiddenElements.forEach((el) => observer.observe(el)); 

const hiddenElements2 = document.querySelectorAll('.about-transition');
hiddenElements2.forEach((el) => observer.observe(el));


const numberSpaces = document.querySelectorAll('.number');
let counters = [0, 0, 0];
let intervals = [];
let percentages = [65, 80, 70];

const circleElement = circle.querySelector('circle');
circleElement.style.setProperty('--dash-offset', dashOffsetVar);

numberSpaces.forEach((number, index) => {
    intervals[index] = setInterval(() => {
        if (counters[index] === percentages[index]) {
            clearInterval(intervals[index]);
        } else {
            counters[index] += 1;
            number.innerHTML = `${counters[index]}%`;
        }
    }, 20);
});