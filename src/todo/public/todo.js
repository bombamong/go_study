(function ($) {
  "use strict";
  $(function () {
    const todoListItem = $(".todo-list");
    const todoListInput = $(".todo-list-input");
    const todoAddButton = $(".todo-list-add-btn");
    const addItem = item =>
      todoListItem.append(
        `<li ${item.completed ? "class=completed" : ""} id=${item.id}>
          <div class='form-check'>
            <label class='form-check-label'>
              <input class='checkbox' type='checkbox' ${
                item.completed ? "checked" : ""
              } />
              ${item.name}
              <i class='input-helper'></i>
            </label>
          </div>
          <i class='remove mdi mdi-close-circle-outline'></i>
        </li>`
      );

    todoAddButton.on("click", function (event) {
      event.preventDefault();
      var item = $(this).prevAll(".todo-list-input").val();
      if (item) {
        $.post("/todos", { name: item }, addItem);
        todoListInput.val("");
      }
    });

    todoListItem.on("change", ".checkbox", function () {
      const id = $(this).closest("li").attr("id");
      const $self = $(this);
      const complete = Boolean($self.attr("checked"));

      $.get(`complete-todo/${id}?complete=${complete}`, function (data) {
        if (complete) {
          $self.attr("checked", "checked");
        } else {
          $self.removeAttr("checked");
        }

        $self.closest("li").toggleClass("completed");
      });
    });

    todoListItem.on("click", ".remove", function () {
      const id = $(this).closest("li").attr("id");
      const $self = $(this);
      console.log(id, $self);
      $.ajax({
        type: "DELETE",
        url: `todos/${id}`,
        success: data => (data.success ? $self.parent().remove() : null),
      });
      $(this).parent().remove();
    });

    $.get("/todos", items => items.forEach(e => addItem(e)));
  });
})(jQuery);
