// familydisplay.js - Display family information

function familydisplay(id) {
    $.getJSON('../api/family/' + id + '.json', function(data) {
        var html = '<div class="family-display">';
        
        // Family title
        html += '<h1>' + (data.Ref.Title || 'Family') + '</h1>';
        
        // Parents
        html += '<div class="parents"><h2>Parents</h2>';
        if (data.Ref.Husband) {
            html += '<div>Husband: ' + createPersonLink(data.Ref.Husband) + '</div>';
        }
        if (data.Ref.Wife) {
            html += '<div>Wife: ' + createPersonLink(data.Ref.Wife) + '</div>';
        }
        html += '</div>';
        
        // Marriage events
        if (data.Events && data.Events.length > 0) {
            html += '<div class="family-events"><h2>Events</h2>';
            data.Events.forEach(function(event) {
                html += displayEvent(event);
            });
            html += '</div>';
        }
        
        // Children
        if (data.Children && data.Children.length > 0) {
            html += '<div class="children"><h2>Children</h2><ul>';
            data.Children.forEach(function(child) {
                html += '<li>' + createPersonLink(child);
                if (child.Birth) {
                    html += ' (b. ' + child.Birth + ')';
                }
                html += '</li>';
            });
            html += '</ul></div>';
        }
        
        // Notes
        if (data.Notes && data.Notes.length > 0) {
            html += '<div class="notes"><h2>Notes</h2>';
            data.Notes.forEach(function(note) {
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
