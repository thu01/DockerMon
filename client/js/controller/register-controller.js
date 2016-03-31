'use strict';

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