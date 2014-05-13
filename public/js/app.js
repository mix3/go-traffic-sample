angular.module('TodoApp', [])
	// config
	.config(function ($interpolateProvider) {
		// {{ }} がgoのテンプレとバッティングするので
		$interpolateProvider.startSymbol('<:');
		$interpolateProvider.endSymbol(':>');
	})

	// controller
	.controller('TodoController', function($scope, $http) {

		// init
		$scope.todos = [];
		$http.get('/list').success(function (data) {
			$scope.todos = data.result
		});

		// for post
		var transformRequest = function(obj) {
			var str = [];
			for (var p in obj) {
				str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
			}
			return str.join("&");
		};

		// create
		$scope.create = function() {
			$http({
				method: 'POST',
				url: '/create',
				transformRequest: transformRequest,
				data: {"title": $scope.newTodoTitle},
				headers: {'Content-Type': 'application/x-www-form-urlencoded'}
			}).success(function(data) {
				$scope.todos = data.result
			})
			$scope.newTodoTitle = '';
		};

		// count
		$scope.getCompletedConut = function() {
			var count = 0;
			angular.forEach($scope.todos, function(todo) {
				count += todo.completed ? 1 : 0;
			});
			return count;
		};

		// delete all
		$scope.deleteAllCompletedTodo = function() {
			$http({
				method: 'POST',
				url: '/delete',
				transformRequest: transformRequest,
				data: {},
				headers: {'Content-Type': 'application/x-www-form-urlencoded'}
			}).success(function(data) {
				$scope.todos = data.result
			})
		};

		// delete
		$scope.deleteCompletedTodo = function(id) {
			$http({
				method: 'POST',
				url: '/delete/' + id,
				transformRequest: transformRequest,
				data: {},
				headers: {'Content-Type': 'application/x-www-form-urlencoded'}
			}).success(function(data) {
				$scope.todos = data.result
			})
		};

		// switch
		$scope.switchCompleted = function(id) {
			$http({
				method: 'POST',
				url: '/switch/' + id,
				transformRequest: transformRequest,
				data: {},
				headers: {'Content-Type': 'application/x-www-form-urlencoded'}
			}).success(function(data) {
				$scope.todos = data.result
			})
		};
	});
