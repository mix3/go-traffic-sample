{{ template "includes/header" }}
	<div ng-controller="TodoController">
		<p>
		Completed Todo: <: getCompletedConut() :> / <: todos.length :>
		<button ng-click="deleteAllCompletedTodo()">Delete Completed Todo</button>
		</p>
		<ul>
			<li ng-repeat="todo in todos">
				<input type="checkbox" ng-model="todo.completed" ng-change="switchCompleted(todo.id)" />
				<span class="completed-<: todo.completed :>"><: todo.title :></span>
				<a href="#" ng-click="deleteCompletedTodo(todo.id)">[x]</a>
			</li>
		</ul>
		<form ng-submit="create()">
			<input type="text" ng-model="newTodoTitle" />
			<input type="submit" value="create" />
		</form>
	</div>
{{ template "includes/footer" }}
