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