// individualdisplay.js - Display individual person information

function individualdisplay(id) {
    $.getJSON('../api/individual/' + id + '.json', function(data) {
        var html = '<div class="individual-display">';
        var ref = data.ref || data.Ref || {};
        var events = data.events || data.Events || [];
        var parents = data.parentsfamily || data.Parents || [];
        var families = data.family || data.Spouses || [];
        var notes = data.notes || data.Notes || [];
        var refName = ref.name || ref.Name || 'Unknown';
        var refPhoto = ref.photo || ref.Photo;
        
        // Name and basic info
        html += '<h1>' + refName + '</h1>';
        
        // Portrait if available
        if (refPhoto) {
            html += '<div class="portrait"><img src="' + refPhoto + '" alt="Portrait"></div>';
        }
        
        // Life events
        if (events.length > 0) {
            html += '<div class="life-events"><h2>Life Events</h2>';
            events.forEach(function(event) {
                html += displayEvent(event);
            });
            html += '</div>';
        }
        
        // Parents
        if (parents.length > 0) {
            html += '<div class="parents"><h2>Parents</h2>';
            parents.forEach(function(parent) {
                html += '<div class="parent-family">';
                var father = parent.father || parent.Husband;
                var mother = parent.mother || parent.Wife;
                if (father) {
                    html += '<div>Father: ' + createPersonLink(father) + '</div>';
                }
                if (mother) {
                    html += '<div>Mother: ' + createPersonLink(mother) + '</div>';
                }
                html += '</div>';
            });
            html += '</div>';
        }
        
        // Spouses and children
        if (families.length > 0) {
            html += '<div class="families"><h2>Family</h2>';
            families.forEach(function(family) {
                html += '<div class="family">';
                var father = family.father || family.Husband;
                var mother = family.mother || family.Wife;
                var fatherID = father && (father.id || father.ID);
                var motherID = mother && (mother.id || mother.ID);

                if (father && fatherID !== id) {
                    html += '<div>Spouse: ' + createPersonLink(father) + '</div>';
                }
                if (mother && motherID !== id) {
                    html += '<div>Spouse: ' + createPersonLink(mother) + '</div>';
                }
                var children = family.children || family.Children || [];
                if (children.length > 0) {
                    html += '<div class="children"><h3>Children</h3><ul>';
                    children.forEach(function(child) {
                        html += '<li>' + createPersonLink(child) + '</li>';
                    });
                    html += '</ul></div>';
                }
                html += '</div>';
            });
            html += '</div>';
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
        $('#display').html('<p class="error">Error loading person data for ID: ' + id + '</p>');
    });
}
