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
    var id = person && (person.id || person.ID);
    if (!id) return 'Unknown';
    var name = (person && (person.name || person.Name)) || 'Unknown';
    return '<a href="../' + id + '/">' + name + '</a>';
}

function displayEvent(event) {
    if (!event) return '';
    var html = '<div class="event">';
    var tag = event.tag || event.Tag || 'Event';
    var date = event.date || event.Date;
    var place = event.place || event.Place;

    html += '<strong>' + tag + '</strong>';
    if (date) {
        html += ' <span class="date">' + formatDate(date) + '</span>';
    }
    if (place) {
        html += ' <span class="place">' + formatPlace(place) + '</span>';
    }
    html += '</div>';
    return html;
}
