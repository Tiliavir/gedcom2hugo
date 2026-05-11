// individualdisplay.js - Display individual person information

function individualdisplay(id) {
    $.getJSON('/api/individual/' + id + '.json', function(data) {
        var html = '<div class="individual-display">';
        
        // Name and basic info
        html += '<h1>' + (data.Ref.Name || 'Unknown') + '</h1>';
        
        // Portrait if available
        if (data.Ref.Photo) {
            html += '<div class="portrait"><img src="' + data.Ref.Photo + '" alt="Portrait"></div>';
        }
        
        // Life events
        if (data.Events && data.Events.length > 0) {
            html += '<div class="life-events"><h2>Life Events</h2>';
            data.Events.forEach(function(event) {
                html += displayEvent(event);
            });
            html += '</div>';
        }
        
        // Parents
        if (data.Parents && data.Parents.length > 0) {
            html += '<div class="parents"><h2>Parents</h2>';
            data.Parents.forEach(function(parent) {
                html += '<div class="parent-family">';
                if (parent.Husband) {
                    html += '<div>Father: ' + createPersonLink(parent.Husband) + '</div>';
                }
                if (parent.Wife) {
                    html += '<div>Mother: ' + createPersonLink(parent.Wife) + '</div>';
                }
                html += '</div>';
            });
            html += '</div>';
        }
        
        // Spouses and children
        if (data.Spouses && data.Spouses.length > 0) {
            html += '<div class="families"><h2>Family</h2>';
            data.Spouses.forEach(function(family) {
                html += '<div class="family">';
                if (family.Husband && family.Husband.ID !== id) {
                    html += '<div>Spouse: ' + createPersonLink(family.Husband) + '</div>';
                }
                if (family.Wife && family.Wife.ID !== id) {
                    html += '<div>Spouse: ' + createPersonLink(family.Wife) + '</div>';
                }
                if (family.Children && family.Children.length > 0) {
                    html += '<div class="children"><h3>Children</h3><ul>';
                    family.Children.forEach(function(child) {
                        html += '<li>' + createPersonLink(child) + '</li>';
                    });
                    html += '</ul></div>';
                }
                html += '</div>';
            });
            html += '</div>';
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
        $('#display').html('<p class="error">Error loading person data for ID: ' + id + '</p>');
    });
}
