'use strict';

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

