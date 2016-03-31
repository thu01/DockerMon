'use strict';

var app = angular.module('app', [
                                    'ui.router',
                                    'ngResource',
                                    'ngMessages'
]);

app.config([
    '$stateProvider',
    '$urlRouterProvider',
    '$locationProvider',
    function($stateProvider, $urlRouterProvider, $locationProvider){
        
        //Remove the # from url
        $locationProvider.html5Mode({
            enabled: true,
            requireBase: false
        });
        
        $stateProvider
        .state('default',{
            url: '/',
            templateUrl: '/html/home.html',
            controller: 'RegisterController'
        })
        .state('login', {
            url: '/login',
            templateUrl: 'html/login.html',
            controller: 'LoginController'
        })
        .state('about',{
            url: '/about',
            templateUrl: '/html/about.html'
        });
        
        $urlRouterProvider.otherwise('/');
}]);

app.factory('UserService', ['$http', '$resource', function UserService($http, $resource) {
    return {
        Register: function(user){
            console.log(user)
            var res = $resource('/api/users', user);
            return res.save().$promise;
            //return $http.post('http://localhost:8000/api/users', user).then(handleSuccess, handleError);
        },
        Login: function(user){
            console.log(user)
            var res = $resource('/api/users', user);
            return res.save().$promise;
            //return $http.post('http://localhost:8000/api/users', user).then(handleSuccess, handleError);
        }
    }
    
    function handleSuccess(res) {
        return res;
    }
    
    function handleError(res) {
        return res;
    }
}]);

app.controller('RegisterController', ['UserService', '$scope', '$location', function(UserService, $scope, $location){
    $scope.user = {}
    $scope.user.username = 'thu';
    $scope.user.email = 'thu@gmail.com';
    $scope.user.password='12345';
    $scope.user.verifyPassword='12345';
    $scope.register = function(user) {
        console.log($scope.registerForm.$pending);
        //If verifyPassword is invalid, the password field will be undefined
        if($scope.registerForm.$pending !== undefined && $scope.registerForm.$pending.username){
            return;
        }
        UserService.Register(user).then(function(response){
            if(response.Status==200) {
                $location.path('/about');
            }
            console.log(response);
            //TODO: error handling
        }, function(response){
            console.log('error');
            console.log(response);
            //TODO: error handling
        });
    };
}]);

app.controller('LoginController', ['UserService','$scope', '$location', function(UserService, $scope, $location){
    $scope.user = {}
    $scope.user.username = 'thu';
    $scope.user.password='12345';
    $scope.login = function(user) {
        UserService.Login(user).then(function(response){
            if(response.Status==200) {
                $location.path('/about');
            }
            console.log(response);
            //TODO: error handling
        }, function(response){
            console.log('error');
            console.log(response);
            //TODO: error handling
        });
    };
}]);


//Custom validation
app.directive('compareTo', function() {
    return {
        require: "ngModel",
        scope: {
            otherModelValue: "=compareTo"
        },
        link: function(scope, elm, attrs, ctrl) {
            ctrl.$validators.compareTo=function(modelValue, viewValue) {
                var res = (modelValue == scope.otherModelValue);
                ctrl.$setValidity('passwordMatch', res); 
                return res;
            }
            scope.$watch("otherModelValue", function() {
                ctrl.$validate();
            });
        }
    };
});

app.directive('username', function($q, $timeout) {
  return {
    require: 'ngModel',
    link: function(scope, elm, attrs, ctrl) {
      var usernames = ['Jim', 'John', 'Jill', 'Jackie'];
      ctrl.$asyncValidators.username = function(modelValue, viewValue) {
        if (ctrl.$isEmpty(modelValue)) {
          // consider empty model valid
          return $q.when();
        }
        var def = $q.defer();
        $timeout(function() {
          // Mock a delayed response
          if (usernames.indexOf(modelValue) === -1) {
            // The username is available
            ctrl.$setValidity('validUsername', true);
            def.resolve();
          } else {
            ctrl.$setValidity('validUsername', false);
            def.reject();
          }

        }, 2000);
        return def.promise;
      };
    }
  };
});