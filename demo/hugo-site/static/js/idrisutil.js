// idrisutil.js - Utility functions for displaying genealogy data

function formatDate(dateStr) {
    if (!dateStr) return '';
    return dateStr;
}

function formatPlace(placeStr) {
    if (!placeStr) return '';
    return placeStr;
}

function createPersonLink(person) {
    if (!person || !person.ID) return 'Unknown';
    var name = person.Name || 'Unknown';
    return '<a href="/' + person.ID + '/">' + name + '</a>';
}

function displayEvent(event) {
    if (!event) return '';
    var html = '<div class="event">';
    html += '<strong>' + (event.Tag || 'Event') + '</strong>';
    if (event.Date) {
        html += ' <span class="date">' + formatDate(event.Date) + '</span>';
    }
    if (event.Place) {
        html += ' <span class="place">' + formatPlace(event.Place) + '</span>';
    }
    html += '</div>';
    return html;
}
