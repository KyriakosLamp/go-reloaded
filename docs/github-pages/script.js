let currentSlideIndex = 1;
let autoAdvanceTimer;

function showSlide(n) {
    const slides = document.querySelectorAll('.carousel-item');
    const dots = document.querySelectorAll('.dot');
    
    if (n > slides.length) {
        currentSlideIndex = 1;
    }
    if (n < 1) {
        currentSlideIndex = slides.length;
    }
    
    // Hide all slides
    slides.forEach(slide => slide.classList.remove('active'));
    dots.forEach(dot => dot.classList.remove('active'));
    
    // Show current slide
    slides[currentSlideIndex - 1].classList.add('active');
    dots[currentSlideIndex - 1].classList.add('active');
}

function changeSlide(n) {
    currentSlideIndex += n;
    showSlide(currentSlideIndex);
}

function currentSlide(n) {
    currentSlideIndex = n;
    showSlide(currentSlideIndex);
    resetAutoAdvance();
}

function resetAutoAdvance() {
    clearInterval(autoAdvanceTimer);
    autoAdvanceTimer = setInterval(() => {
        changeSlide(1);
    }, 5000);
}

// Section scrolling
let isScrolling = false;
let scrollTimeout;

function scrollToSection(direction) {
    if (isScrolling) return;
    
    const sections = document.querySelectorAll('.section');
    const currentSection = getCurrentSection();
    let targetIndex = currentSection;
    
    if (direction > 0 && targetIndex < sections.length - 1) {
        targetIndex++;
    } else if (direction < 0 && targetIndex > 0) {
        targetIndex--;
    }
    
    if (targetIndex !== currentSection) {
        isScrolling = true;
        sections[targetIndex].scrollIntoView({ behavior: 'smooth' });
        
        setTimeout(() => {
            isScrolling = false;
        }, 1000);
    }
}

function getCurrentSection() {
    const sections = document.querySelectorAll('.section');
    let currentIndex = 0;
    
    sections.forEach((section, index) => {
        const rect = section.getBoundingClientRect();
        if (rect.top <= window.innerHeight / 2 && rect.bottom >= window.innerHeight / 2) {
            currentIndex = index;
        }
    });
    
    return currentIndex;
}

// Wheel event for section scrolling
window.addEventListener('wheel', (e) => {
    e.preventDefault();
    
    clearTimeout(scrollTimeout);
    scrollTimeout = setTimeout(() => {
        if (e.deltaY > 0) {
            scrollToSection(1); // Scroll down
        } else {
            scrollToSection(-1); // Scroll up
        }
    }, 50);
}, { passive: false });

// Initialize carousel and navigation
document.addEventListener('DOMContentLoaded', () => {
    showSlide(currentSlideIndex);
    updateActiveSection();
    resetAutoAdvance();
});

// Section navigation
function updateActiveSection() {
    const sections = document.querySelectorAll('.section');
    const navLinks = document.querySelectorAll('.nav-link');
    
    let currentSection = '';
    
    sections.forEach(section => {
        const rect = section.getBoundingClientRect();
        if (rect.top <= window.innerHeight / 2 && rect.bottom >= window.innerHeight / 2) {
            currentSection = section.id;
        }
    });
    
    navLinks.forEach(link => {
        link.classList.remove('active');
        if (link.getAttribute('href') === '#' + currentSection) {
            link.classList.add('active');
        }
    });
}

// Update active section on scroll
window.addEventListener('scroll', updateActiveSection);

// Smooth scrolling for navigation links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth'
            });
        }
    });
});

