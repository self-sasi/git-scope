// ===============================
// ANALYTICS SETUP
// ===============================
const analytics = {
    events: [],
    track: function (name, data = {}) {
        const event = { name, data, time: new Date().toISOString() };
        this.events.push(event);
        console.log('[Analytics]', name, data);
    }
};

analytics.track('page_view', { page: window.location.pathname });

// ===============================
// GOATCOUNTER EVENT TRACKING
// ===============================
function trackEvent(eventName, title) {
    if (window.goatcounter && window.goatcounter.count) {
        window.goatcounter.count({
            path: eventName,
            title: title || eventName,
            event: true
        });
    }
}

// ===============================
// SCROLL DEPTH TRACKING
// ===============================
let maxScroll = 0;
let scrollMilestones = { 25: false, 50: false, 75: false, 100: false };
window.addEventListener('scroll', () => {
    const scrollPercent = Math.round((window.scrollY / (document.body.scrollHeight - window.innerHeight)) * 100);
    if (scrollPercent > maxScroll) {
        maxScroll = scrollPercent;
        [25, 50, 75, 100].forEach(milestone => {
            if (scrollPercent >= milestone && !scrollMilestones[milestone]) {
                scrollMilestones[milestone] = true;
                trackEvent('scroll-' + milestone + '-percent', 'Scrolled ' + milestone + '% of page');
            }
        });
    }
});

// ===============================
// STAR COUNT FETCHER
// ===============================
fetch('https://api.github.com/repos/Bharath-code/git-scope')
    .then(response => response.json())
    .then(data => {
        if (data.stargazers_count !== undefined) {
            const count = data.stargazers_count;
            const formatted = count > 1000 ? (count / 1000).toFixed(1) + 'k' : count;
            document.getElementById('star-count').textContent = `Star ${formatted}`;
        }
    })
    .catch(err => console.log('Failed to fetch stars', err));

// ===============================
// TAB SWITCHING
// ===============================
function switchTab(id) {
    document.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));
    event.target.classList.add('active');
    document.querySelectorAll('.code-block').forEach(b => b.classList.remove('active'));
    document.getElementById('code-' + id).classList.add('active');
    trackEvent('tab-' + id, 'Switched to ' + id + ' install tab');
}

// ===============================
// COPY COMMAND
// ===============================
function copyCode() {
    const activeBlock = document.querySelector('.code-block.active');
    let text = "";
    activeBlock.querySelectorAll('.cmd').forEach(el => text += el.innerText + "\n");
    text = text.trim();

    navigator.clipboard.writeText(text).then(() => {
        const hint = document.querySelector('.copy-hint');
        const original = hint.innerText;
        hint.innerText = "Copied!";
        hint.style.opacity = '1';
        setTimeout(() => {
            hint.innerText = original;
        }, 2000);
        trackEvent('command-copied', 'Copied install command');
    });
}

// ===============================
// GITHUB BUTTON CLICK TRACKING
// ===============================
document.getElementById('github-btn')?.addEventListener('click', () => {
    trackEvent('github-click-nav', 'Clicked GitHub button in nav');
});

document.querySelectorAll('footer a').forEach(link => {
    link.addEventListener('click', () => {
        const linkName = link.textContent.trim();
        trackEvent('footer-click-' + linkName.toLowerCase(), 'Clicked ' + linkName + ' in footer');
    });
});

// ===============================
// SESSION SUMMARY
// ===============================
window.addEventListener('beforeunload', () => {
    analytics.track('session_end', {
        maxScroll: maxScroll,
        totalEvents: analytics.events.length
    });
});
