// familydisplay.js - Display family information

function familydisplay(id) {
    $.getJSON('../api/family/' + id + '.json', function(data) {
        var html = '<div class="family-display">';
        var ref = data.ref || data.Ref || {};
        var events = data.events || data.Events || [];
        var children = data.children || data.Children || [];
        var notes = data.note || data.Notes || [];
        var husband = ref.husband || ref.Husband;
        var wife = ref.wife || ref.Wife;
        var title = ref.title || ref.Title || 'Family';
        
        // Family title
        html += '<h1>' + title + '</h1>';
        
        // Parents
        html += '<div class="parents"><h2>Parents</h2>';
        if (husband) {
            html += '<div>Husband: ' + createPersonLink(husband) + '</div>';
        }
        if (wife) {
            html += '<div>Wife: ' + createPersonLink(wife) + '</div>';
        }
        html += '</div>';
        
        // Marriage events
        if (events.length > 0) {
            html += '<div class="family-events"><h2>Events</h2>';
            events.forEach(function(event) {
                html += displayEvent(event);
            });
            html += '</div>';
        }
        
        // Children
        if (children.length > 0) {
            html += '<div class="children"><h2>Children</h2><ul>';
            children.forEach(function(child) {
                html += '<li>' + createPersonLink(child);
                var birth = child.birth || child.Birth;
                if (birth) {
                    html += ' (b. ' + birth + ')';
                }
                html += '</li>';
            });
            html += '</ul></div>';
        }
        
        // Notes
        if (notes.length > 0) {
            html += '<div class="notes"><h2>Notes</h2>';
            notes.forEach(function(note) {
                html += '<p>' + note + '</p>';
            });
            html += '</div>';
        }
        
        html += '</div>';
        $('#display').html(html);
        
        // Show raw data for debugging
        $('#raw').html('<details><summary>Raw JSON Data</summary><pre>' + 
                      JSON.stringify(data, null, 2) + '</pre></details>');
    }).fail(function() {
        $('#display').html('<p class="error">Error loading family data for ID: ' + id + '</p>');
    });
}
