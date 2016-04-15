'use strict';

app.controller(
    'LoginController', 
    ['UserService','$scope', '$state', '$cookies',
    function(UserService, $scope, $state, $cookies){
    $scope.user = {}
    $scope.user.username = 'thu';
    $scope.user.password='12345';
    $scope.login = function(user) {
        UserService.Login(user).then(function(response){
            console.log(response);
            if(response.status==200) {
                $state.go('about');
            }
            //TODO: error handling
            console.log('error');
        }, function(response){
            console.log('error');
            console.log(response);
            //TODO: error handling
        });
    };
}]);

