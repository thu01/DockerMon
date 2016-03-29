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
    $scope.user.password='123';
    $scope.user.verifyPassword='123';
    $scope.register = function(user) {
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